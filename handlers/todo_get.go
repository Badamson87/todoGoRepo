package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
	"database/sql"
)

// func TodoGet(todoModel todo.Getter) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         results := todoModel.GetAll()
//         c.JSON(http.StatusOK, results)
//     }
// }

func TodoGet(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        results := todo.GetAll(db)
        c.JSON(http.StatusOK, results)
    }
}


