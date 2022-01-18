package data

import (
	"context"

	v1 "helloworld2/api/helloworld/v1"
	"helloworld2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	r.log.Debugf("CreateGreeter: %v", g)
	r.data.hgclient.SayHello(ctx, &v1.HelloRequest{Name: g.Hello + "from http"})
	r.data.gclient.SayHello(ctx, &v1.HelloRequest{Name: g.Hello + " from grpc"})
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
