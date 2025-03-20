package main

import (
	"flag"
	"fmt"

	"go-service-pb/gorpc"
	"go-service-pb/internal/config"
	gogameclientServer "go-service-pb/internal/server/gogameclient"
	gopmpclientServer "go-service-pb/internal/server/gopmpclient"
	gosysclientServer "go-service-pb/internal/server/gosysclient"
	gouserclientServer "go-service-pb/internal/server/gouserclient"
	govpnclientServer "go-service-pb/internal/server/govpnclient"
	"go-service-pb/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gorpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		gorpc.RegisterGoGameClientServer(grpcServer, gogameclientServer.NewGoGameClientServer(ctx))
		gorpc.RegisterGoPmpClientServer(grpcServer, gopmpclientServer.NewGoPmpClientServer(ctx))
		gorpc.RegisterGoSysClientServer(grpcServer, gosysclientServer.NewGoSysClientServer(ctx))
		gorpc.RegisterGoUserClientServer(grpcServer, gouserclientServer.NewGoUserClientServer(ctx))
		gorpc.RegisterGoVpnClientServer(grpcServer, govpnclientServer.NewGoVpnClientServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
