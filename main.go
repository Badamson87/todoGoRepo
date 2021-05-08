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
    api := r.Group("/api")
    {
        api.GET("/todo", handlers.TodoGet(todoModel))
        api.POST("/todo", handlers.TodoPost(todoModel))
        api.DELETE("/todo", handlers.TodoPost(todoModel))
        api.PUT("/todo", handlers.TodoPost(todoModel))
    }
    r.Run("0.0.0.0:5000")
  }

