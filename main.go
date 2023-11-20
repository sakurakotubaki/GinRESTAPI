package main

import(
  "github.com/gin-gonic/gin"
  "net/http"
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
  r.Run()
}