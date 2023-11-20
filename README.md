# Ginを使用して、REST APIを作成する
[Ginとは何か?](https://gin-gonic.com/ja/docs/)

1. プロジェクトを作成する
```bash
mkdir gin_api && cd gin_api
```
2. モジュールを初期化する
```bash
go mod init gin_api
```
3. Ginをインストールする
```bash
go get -u github.com/gin-gonic/gin
```
4. main.goを作成する
```bash
touch main.go
```

Hello Worldを表示するコードを書く
```go
package main

import(
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run()
}
```

5. main.goを実行する
```bash
go run main.go
```
6. curlでport8080にアクセスする。 -X GETを指定することで、GETリクエストを送信することができる。省略は可能。
```bash
curl http://localhost:8080 -X GET
```

7. ダミーのデータを返したり、保存する配列を作成して、GET POSTに対応する
```go
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
```

8. curlでPOSTリクエストを送信する
```bash
curl http://localhost:8080/todos -X POST -H "Content-Type: application/json" -d '{"id": 1, "text": "test"}'
```

9. curlでGETリクエストを送信する
```bash
curl http://localhost:8080/todos -X GET
```