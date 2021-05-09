package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo/models"
	"database/sql"
)

func TodoDelete(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
    var ids = c.Query("id");
        results := todo.Delete(ids, db)
        c.JSON(http.StatusOK, results)
    }
}

func TodoSoftDelete(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
    var ids = c.Query("id");
        results := todo.SoftDelete(ids, db)
        c.JSON(http.StatusOK, results)
    }
}


