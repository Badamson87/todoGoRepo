package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "todo/handlers"
    "database/sql"
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
        api.POST("/todo", handlers.TodoPost(db))
//         api.DELETE("/todo", handlers.TodoPost(todoModel))
//         api.PUT("/todo", handlers.TodoPost(todoModel))
    }

//     insert, err := db.Query("INSERT INTO todo (checked, title) VALUES (0, 'dbTest')")
//     if err != nil {
//         panic(err.Error())
//     }
//     defer insert.Close()

//     fmt.Println("inserted user")
//
    r.Run("0.0.0.0:5000")
  }
