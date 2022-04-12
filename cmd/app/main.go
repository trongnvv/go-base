package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-base/pkg/crud/fxloader"
	grpc_base "go-base/pkg/crud/grpc"
	pb "go-base/pkg/crud/grpc/define"
	"go-base/pkg/helpers/configs"
	"go-base/pkg/helpers/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func init() {
	configs.LoadConfig("configs/config.json")
	log.LoadLogger(configs.Get())
}

func main() {
	fx.New(
		fx.Provide(configs.Get),
		fx.Provide(log.GetZeroLog),
		fx.Options(fxloader.LoadFX()...),
		fx.Invoke(fxHooks),
	).Run()
}

func fxHooks(lc fx.Lifecycle, exampleService *grpc_base.ExampleService, ginEngine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				lis, err := net.Listen("tcp", configs.Get().Address.Server.Internal)
				if err != nil {
					log.Fatalf(err, "cannot listen, address:[%s]", configs.Get().Address.Server.Internal)
				}
				grpcServer := grpc.NewServer()
				pb.RegisterExampleServer(grpcServer, exampleService)
				if err = grpcServer.Serve(lis); err != nil {
					log.Fatal(err, "cannot serve")
				}
			}()
			go func() {
				if err := ginEngine.Run(configs.Get().Address.Server.Restful); err != nil {
					log.Fatal(err, "Cannot start application gin engine")
				} else {
					log.Info("http restful is running")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping restful server.")
			log.Info("Stopping payment integrator server.")
			return nil
		},
	})
}
