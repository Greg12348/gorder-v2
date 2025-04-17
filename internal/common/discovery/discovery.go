package discovery

import (
	"context"
	"fmt"
	"github.com/Greg12348/gorder-v2/common/discovery/consul"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPost string) error
	Deregister(ctx context.Context, instanceID, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceID, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	x := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	return fmt.Sprintf("%s-%d", serviceName, x)
}
func GetServiceAddr(ctx context.Context, serviceName string) (string, error) {
	registry, err := consul.New(viper.GetString("consul.addr"))
	if err != nil {
		return "", err
	}
	addrs, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("service %s not found in consul", serviceName)
	}
	i := rand.Intn(len(addrs))
	logrus.Infof("discovered %d service %s at %s", len(addrs), serviceName, addrs[i])
	return addrs[i], nil
}
