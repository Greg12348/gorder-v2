package main

import (
	"context"
	"github.com/Greg12348/gorder-v2/common/config"
	"github.com/Greg12348/gorder-v2/common/genproto/stockpb"
	"github.com/Greg12348/gorder-v2/common/server"
	"github.com/Greg12348/gorder-v2/stock/ports"
	"github.com/Greg12348/gorder-v2/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.service-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// not use
	default:
		panic("unexpected server type")
	}
}
