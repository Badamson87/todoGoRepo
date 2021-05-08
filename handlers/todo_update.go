package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
)

type todoPutRequest struct {
    Id int `json:"id"`
    Checked bool  `json:"checked"`
    Title string  `json:"title"`
}

func TodoUpdate(todoModel todo.Updater) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := todoPutRequest{}
        c.Bind(&requestBody)
            todo := todo.Item{
                Id: requestBody.Id,
                Checked: requestBody.Checked,
                Title: requestBody.Title,
            }
        todoModel.Update(todo)
        c.Status(http.StatusNoContent)
    }
}


