package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"io"
	"strong-defi-backend/pkg/log"
	"time"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String() // Generate a new trace_id
		}
		log.SetTraceId(traceID)
		log.SetKeyValue("URI", c.Request.RequestURI)
		c.Set("trace_id", traceID)

		c.Next()

		rq, _ := io.ReadAll(c.Request.Body)
		log.Info().
			Dict("request", zerolog.Dict().
				Str("uri", c.Request.RequestURI).
				Str("requestData", string(rq)),
			).Msgf("duration_time: %d", float32((time.Now().UnixMilli()-t)/1000))
	}
}
