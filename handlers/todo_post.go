package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
)

type todoPostRequest struct {
    Id int `json:"id"`
    Checked bool  `json:"checked"`
    Title string  `json:"title"`
}

func TodoPost(todoModel todo.Adder) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := todoPostRequest{}
        c.Bind(&requestBody)

            todo := todo.Item{
                Id: requestBody.Id,
                Checked: requestBody.Checked,
                Title: requestBody.Title,
            }
        todoModel.Add(todo)
        c.Status(http.StatusNoContent)
    }
}


