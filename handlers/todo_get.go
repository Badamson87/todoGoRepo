package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
)

func TodoGet(todoModel todo.Getter) gin.HandlerFunc {
    return func(c *gin.Context) {
        results := todoModel.GetAll()
        c.JSON(http.StatusOK, results)
    }
}


