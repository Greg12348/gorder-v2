package client

import (
	"context"
	"github.com/Greg12348/gorder-v2/common/discovery"
	"github.com/Greg12348/gorder-v2/common/genproto/stockpb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStockGRPCClient(ctx context.Context) (client stockpb.StockServiceClient, close func() error, err error) {
	//grpcAddr := viper.GetString("stock.grpc-addr")
	grpcAddr, err := discovery.GetServiceAddr(ctx, viper.GetString("stock.service-name"))
	if err != nil {
		return nil, func() error {
			return nil
		}, err
	}
	if grpcAddr == "" {
		logrus.Warn("no stock service address found")
	}
	opts, err := grpcDialOptions(grpcAddr)
	if err != nil {
		return nil, func() error {
			return nil
		}, err
	}
	conn, err := grpc.NewClient(grpcAddr, opts...)
	if err != nil {
		return nil, func() error {
			return nil
		}, err
	}
	return stockpb.NewStockServiceClient(conn), conn.Close, nil
}

func grpcDialOptions(addr string) ([]grpc.DialOption, error) {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
}
