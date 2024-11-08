package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strong-defi-backend/pkg/log"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String() // Generate a new trace_id
		}
		log.SetTraceId(traceID)
		c.Set("trace_id", traceID)
		log.Info().Msgf("xxxxxxxxxxxxxxx set::")
		c.Next()
		log.Info().Msgf("xxxxxxxxxxxxxxx get::")
	}
}
