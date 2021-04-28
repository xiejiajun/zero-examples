// Code generated by goctl. DO NOT EDIT!
// Source: add.proto

package main

import (
	"flag"
	"fmt"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/config"
	"bookstore/rpc/add/internal/server"
	"bookstore/rpc/add/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/add.yaml", "the config file")

// TODO 多rpc接口案例
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	adderSrv := server.NewAdderServer(ctx)

	s, err := zrpc.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAdderServer(grpcServer, adderSrv)
	})
	logx.Must(err)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
