package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "todo/handlers"
    "database/sql"
//     "net/http"
    _ "github.com/go-sql-driver/mysql"

    )

func main() {
    // establish DB connection
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tododb")
    if err != nil{
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("connected to mysql")

    r := gin.Default()
    api := r.Group("/api")
    {
        api.GET("/todo", handlers.TodoGet(db))
        api.GET("/todo/history", handlers.TodoGetHistory(db))
        api.POST("/todo", handlers.TodoPost(db))
        api.PUT("/todo", handlers.TodoUpdate(db))
        api.DELETE("/todo/" , handlers.TodoDelete(db))
        api.DELETE("/todo/soft/" , handlers.TodoSoftDelete(db))
    }
//      r.Static("/", "public/dist/typescriptAngular")
//  r.GET("/hello", func(c *gin.Context) {
//         c.String(http.StatusOK, "Hello from %v", "Gin")
//     })
//         r.GET("", func(c *gin.Context) {
//             c.HTML(http.StatusOK, "public/dist/typescriptAngular/index.html", nil)
//         })

//         r.NoRoute(func(c *gin.Context) {
//                 c.File("public/dist/typescriptAngular/index.html")
//             })
//
//     r.GET("/", func(c *gin.Context) {
//     	 c.HTML(http.StatusOK, "public/dist/typescriptAngular/index.html", gin.H{
//     	        "title": "Main website",
//     	    })
//     	})
    r.Run("0.0.0.0:5000")
  }
