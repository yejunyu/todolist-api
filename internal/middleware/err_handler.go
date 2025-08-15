package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
				if !context.Writer.Written() {
					context.JSON(http.StatusInternalServerError, gin.H{
						"error": "Internal Server Error",
					})
				}
			}
		}()
	}
}
