package main

import (
	"flag"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/zeromicro/zero-examples/tracing/remote/portal"
)

var (
	configFile = flag.String("f", "config.json", "the config file")
	client     zrpc.Client
)

type Config struct {
	rest.RestConf
	Portal zrpc.RpcClientConf
}

func handle(w http.ResponseWriter, r *http.Request) {
	conn := client.Conn()
	greet := portal.NewPortalClient(conn)
	resp, err := greet.Portal(r.Context(), &portal.PortalRequest{
		Name: "kevin",
	})
	if err != nil {
		httpx.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	} else {
		httpx.OkJson(w, resp.Response)
	}
}

// TODO http服务
func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)
	client = zrpc.MustNewClient(c.Portal)
	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: c.Port,
	})
	defer engine.Stop()

	engine.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: handle,
	})
	engine.Start()
}
