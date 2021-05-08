package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "todo/handlers"
    "todo/models"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    )

var DB *sql.DB

func main() {
    // db connection would go here and pass the db
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tododb")
    if err != nil{
        panic(err.Error())
    }
    DB = db
    defer db.Close()

    fmt.Println("connected to mysql")

    todoModel := todo.New()
    r := gin.Default()
    api := r.Group("/api")
    {
        api.GET("/todo", handlers.TodoGet(db))
        api.POST("/todo", handlers.TodoPost(todoModel))
        api.DELETE("/todo", handlers.TodoPost(todoModel))
        api.PUT("/todo", handlers.TodoPost(todoModel))
    }

    insert, err := db.Query("INSERT INTO todo (checked, title) VALUES (0, 'dbTest')")
    if err != nil {
        panic(err.Error())
    }
    defer insert.Close()

    fmt.Println("inserted user")

    r.Run("0.0.0.0:5000")
  }

