package fxloader

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-base/pkg/crud/adapters/pgsql"
	"go-base/pkg/crud/adapters/pgsql/repositories"
	"go-base/pkg/crud/core/usecases"
	"go-base/pkg/crud/grpc"
	"go-base/pkg/crud/restful/controllers"
	"go-base/pkg/crud/restful/middlewares"
	"go-base/pkg/crud/restful/routers"
	"go-base/pkg/helpers/gin_custom"
	"go.uber.org/fx"
)

func LoadFX() []fx.Option {
	return []fx.Option{
		fx.Options(loadController()...),
		fx.Options(loadMiddleware()...),
		fx.Options(loadGinEngine()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadGrpcService()...),
		fx.Options(loadAdapterRepo()...),
	}
}

func loadController() []fx.Option {
	return []fx.Option{
		fx.Provide(controllers.NewExampleController),
	}
}

func loadMiddleware() []fx.Option {
	return []fx.Option{
		fx.Provide(middlewares.NewExampleMiddleware),
	}
}

func loadGinEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(gin.New),
		fx.Provide(validator.New),
		fx.Provide(gin_custom.NewBaseController),
		fx.Provide(gin_custom.NewCustomMiddleware),
		fx.Invoke(gin_custom.RegisterCustomMiddleware),
		fx.Invoke(routers.RegisterGinRouters),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(usecases.NewExampleUseCase),
	}
}

func loadGrpcService() []fx.Option {
	return []fx.Option{
		fx.Provide(grpc.NewExampleService),
	}
}

func loadAdapterRepo() []fx.Option {
	return []fx.Option{
		fx.Provide(pgsql.NewGormDB),
		fx.Provide(repositories.NewExampleRepository),
	}
}
