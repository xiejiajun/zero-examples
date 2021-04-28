package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
)

var configFile = flag.String("f", "config.json", "the config file")

// TODO grpc普通接口调用测试
func main() {
	flag.Parse()

	var c zrpc.RpcClientConf
	conf.MustLoad(*configFile, &c)
	client := zrpc.MustNewClient(c)
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn := client.Conn()
			greet := unary.NewGreeterClient(conn)
			resp, err := greet.Greet(context.Background(), &unary.Request{
				Name: "kevin",
			})
			if err != nil {
				fmt.Println("X", err.Error())
			} else {
				fmt.Println("=>", resp.Greet)
			}
		}
	}
}
