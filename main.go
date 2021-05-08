package main

import (
    "github.com/gin-gonic/gin"
    "todo/handlers"
    "todo/models"
    )


func main() {
    // db connection would go here and pass the db
    todoModel := todo.New()
    r := gin.Default()

    r.GET("/todo", handlers.TodoGet(todoModel))
    r.POST("/todo", handlers.TodoPost(todoModel))

    r.Run() //  listen and serve on :8080
  }

