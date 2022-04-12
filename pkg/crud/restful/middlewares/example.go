package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-base/pkg/helpers/log"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) JustNext(c *gin.Context) {
	log.Info("Pass example middleware")
	c.Next()
}
