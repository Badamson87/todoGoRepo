package handlers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
	"database/sql"
)

type todoPostRequest struct {
    Id int `json:"id"`
    Checked bool  `json:"checked"`
    Title string  `json:"title"`
}

func TodoPost(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := todoPostRequest{}
        fmt.Println(requestBody)
        c.Bind(&requestBody)
            item := todo.Item {
                Id: requestBody.Id,
                Checked: requestBody.Checked,
                Title: requestBody.Title,
            }
        fmt.Println(item)
        results := todo.Add(item, db)
        c.JSON(http.StatusOK, results)
    }
}


