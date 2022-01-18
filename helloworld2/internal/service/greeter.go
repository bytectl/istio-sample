package service

import (
	"context"

	v1 "helloworld2/api/helloworld2/v1"
	"helloworld2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeter2Server

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello2(ctx context.Context, in *v1.Hello2Request) (*v1.Hello2Reply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	s.uc.Create(ctx, &biz.Greeter{Hello: in.GetName()})
	return &v1.Hello2Reply{Message: "server2 Hello " + in.GetName()}, nil
}
