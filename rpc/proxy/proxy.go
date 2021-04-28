package main

import (
	"context"
	"flag"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
	"google.golang.org/grpc"
)

var (
	listen = flag.String("listen", "0.0.0.0:3456", "the address to listen on")
	server = flag.String("server", "dns:///unaryserver:3456", "the backend service")
)

type GreetServer struct {
	*zrpc.RpcProxy
}

func (s *GreetServer) Greet(ctx context.Context, req *unary.Request) (*unary.Response, error) {
	conn, err := s.TakeConn(ctx)
	if err != nil {
		return nil, err
	}

	remote := unary.NewGreeterClient(conn)
	return remote.Greet(ctx, req)
}

// TODO Rpc代理测试
func main() {
	flag.Parse()

	proxy := zrpc.MustNewServer(zrpc.RpcServerConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		ListenOn: *listen,
	}, func(grpcServer *grpc.Server) {
		unary.RegisterGreeterServer(grpcServer, &GreetServer{
			RpcProxy: zrpc.NewProxy(*server),
		})
	})
	proxy.Start()
}
