package main

import(
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
)

type Todo struct {
  ID int `json:"id"`
  Text string `json:"text"`
}

var todos []Todo

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello, World!",
    })
  })
  r.GET("/todos", func(c *gin.Context) {
    if len(todos) == 0 {
      c.JSON(http.StatusNotFound, gin.H{"status": "Not Found"})
      return
    }
    c.JSON(http.StatusOK, todos)
  })
  r.POST("/todos", func(c *gin.Context) {
    var todo Todo
    if err := c.BindJSON(&todo); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    todos = append(todos, todo)
    c.JSON(http.StatusOK, todos)
  })
  r.PUT("/todos/:id", func(c *gin.Context) {
    var newTodo Todo
    if err := c.BindJSON(&newTodo); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
      return
    }
    for i, t := range todos {
      if t.ID == id {
        todos[i] = newTodo
        c.JSON(http.StatusOK, todos)
        return
      }
    }
    c.JSON(http.StatusNotFound, gin.H{"status": "Not Found"})
  })
  r.DELETE("/todos/:id", func(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
      return
    }
    for i, t := range todos {
      if t.ID == id {
        todos = append(todos[:i], todos[i+1:]...)
        c.JSON(http.StatusOK, todos)
        return
      }
    }
    c.JSON(http.StatusNotFound, gin.H{"status": "Not Found"})
  })
  r.Run()
}