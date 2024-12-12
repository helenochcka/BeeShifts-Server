package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-Id")
		if id == "" {
			id = uuid.NewString()
		}

		c.Set("request_id", id)

		c.Next()
	}
}

func Logging(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("request_id")
		if exists != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "request id is missing in context"})
			c.Abort()
			return
		}
		requestLogger := logger.With(slog.String("request_id", id.(string)))
		slog.SetDefault(requestLogger)

		c.Next()
	}
}
