package main

import (
	"time"

	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kitetcd "github.com/go-kit/kit/sd/etcd"
)

// Register 注册服务实例
func Register(ctx context.Context, logger log.Logger, etcdAddrs []string, prefix, instance, accessAddr string) sd.Registrar {
	client, err := kitetcd.NewClient(ctx, etcdAddrs, kitetcd.ClientOptions{})
	if err != nil {
		logger.Log("Etcd Client", err)
		panic(err)
	}
	key := prefix + instance
	value := accessAddr

	registrator := kitetcd.NewRegistrar(client, kitetcd.Service{
		Key:   key,
		Value: value,
		TTL:   kitetcd.NewTTLOption(time.Second*3, time.Second*10),
	}, logger)

	// registrator.Register()

	return registrator
}
