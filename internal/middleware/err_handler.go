package middleware

import (
	"errors"
	"log"
	"net/http"
	"todolist-api/pkg/ierr"
	"todolist-api/pkg/response"

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
		context.Next()
		if len(context.Errors) > 0 {
			err := context.Errors.Last().Err
			var apiErr *ierr.APIError
			if errors.As(err, &apiErr) {
				response.Fail(context, apiErr.Msg)
				return
			}
			// 如果不是自定义的错误
			response.Fail(context, ierr.ErrSystem.Msg)
		}
	}
}
