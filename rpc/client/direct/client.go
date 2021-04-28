package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
)

const timeFormat = "15:04:05"

// TODO gRpc服务发现方式调用测试
func main() {
	flag.Parse()

	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"localhost:2379"},
			Key:   "zrpc",
		},
	})

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn := client.Conn()
			greet := unary.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			resp, err := greet.Greet(ctx, &unary.Request{
				Name: "kevin",
			})
			if err != nil {
				fmt.Printf("%s X %s\n", time.Now().Format(timeFormat), err.Error())
			} else {
				fmt.Printf("%s => %s\n", time.Now().Format(timeFormat), resp.Greet)
			}
			cancel()
		}
	}
}
