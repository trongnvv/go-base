package gin_custom

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"
)

type CustomMiddleware struct {
}

func NewCustomMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}

func RegisterCustomMiddleware(middleware *CustomMiddleware, engine *gin.Engine, logger *zerolog.Logger) {
	engine.Use(middleware.WrapZeroLog(logger))
	engine.Use(middleware.Recovery(logger))
}

func (m *CustomMiddleware) WrapZeroLog(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		logger.Info().
			Str("path", path).
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("query", c.Request.URL.RawQuery).
			Str("user_agent", c.Request.UserAgent()).
			Msg("")
		c.Next()
	}
}

func (m *CustomMiddleware) Recovery(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				logger.Error().
					Str("message", "[Recovery from panic]").
					Time("time", time.Now()).
					Interface("error", err).
					Str("request", string(httpRequest)).
					Str("stack", string(debug.Stack())).
					Msg("")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
