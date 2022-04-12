package routers

import (
	"github.com/gin-gonic/gin"
	"go-base/pkg/crud/restful/controllers"
	"go-base/pkg/crud/restful/middlewares"
	"go-base/pkg/helpers/log"
	"go.uber.org/fx"
	"net/http"
)

type RouterIn struct {
	fx.In
	Engine            *gin.Engine
	ExampleController *controllers.ExampleController
	ExampleMiddleware *middlewares.ExampleMiddleware
}

func RegisterGinRouters(in RouterIn) {
	group := in.Engine.Group("/")
	publicRouter(group, in)
	group.Use(in.ExampleMiddleware.JustNext)
	{
		protectedRouter(group, in)
	}
}

func publicRouter(r *gin.RouterGroup, in RouterIn) {
	r.GET("/ping", func(c *gin.Context) {
		log.Info(
			"ping ping")
		c.String(http.StatusOK, "pong")
	})
	r.POST("/example", in.ExampleController.Create)
	r.GET("/example", in.ExampleController.Read)
	r.PUT("/example/:id", in.ExampleController.Update)
	r.DELETE("/example/:id", in.ExampleController.Delete)
}

func protectedRouter(r *gin.RouterGroup, in RouterIn) {
	r.GET("/user/info", func(c *gin.Context) {
		c.String(http.StatusOK, "example protected api")
	})
}
