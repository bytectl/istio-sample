package data

import (
	"context"
	"helloworld2/internal/conf"

	v1 "helloworld2/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewGreeterHttpClient, NewGreeterClient)

// Data .
type Data struct {
	gclient  v1.GreeterClient
	hgclient v1.GreeterHTTPClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, gclient v1.GreeterClient, hgclient v1.GreeterHTTPClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		gclient:  gclient,
		hgclient: hgclient,
	}, cleanup, nil
}

func NewGreeterClient() v1.GreeterClient {
	// direct:///helloworld1-svc:9000
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("direct:///helloworld1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewGreeterClient(conn)
}

func NewGreeterHttpClient() v1.GreeterHTTPClient {
	// http: //helloworld1-svc:8000?isSecure=false
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("http://helloworld1:8000?isSecure=false"),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewGreeterHTTPClient(conn)
}
