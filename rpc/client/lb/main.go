package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
)

var lb = flag.String("t", "direct", "the load balancer type")

// TODO 负载均衡调用测试，direct方式需要配置多个gRpc Server endpoint
//  etcd服务发现模式只需要指定etcd地址即可
func main() {
	flag.Parse()

	var cli zrpc.Client
	switch *lb {
	case "direct":
		cli = zrpc.MustNewClient(zrpc.RpcClientConf{
			Endpoints: []string{
				"localhost:3456",
				"localhost:3457",
			},
		})
	case "discov":
		cli = zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{
				Hosts: []string{"localhost:2379"},
				Key:   "zrpc",
			},
		})
	default:
		log.Fatal("bad load balancing type")
	}

	greet := unary.NewGreeterClient(cli.Conn())
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
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
