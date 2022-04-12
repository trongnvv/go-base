package grpc

import (
	"context"
	"go-base/pkg/crud/core/entities"
	"go-base/pkg/crud/core/usecases"
	pb "go-base/pkg/crud/grpc/define"
	"go-base/pkg/helpers/log"
)

type ExampleService struct {
	pb.UnimplementedExampleServer
	exampleUseCase *usecases.ExampleUseCase
}

func NewExampleService(exampleUseCase *usecases.ExampleUseCase) *ExampleService {
	return &ExampleService{exampleUseCase: exampleUseCase}
}

func (s *ExampleService) Create(c context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	err := s.exampleUseCase.Create(&entities.Example{Name: req.Name})
	if err != nil {
		log.Errorf(err, "Create repositories fail")
		return nil, err
	}
	return nil, nil
}
