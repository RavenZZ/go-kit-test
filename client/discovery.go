package main

import (
	"context"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"

	"github.com/go-kit/kit/log"
	kitetcd "github.com/go-kit/kit/sd/etcd"
)

var (
	instancer *kitetcd.Instancer = nil
)

func getInstancer(ctx context.Context, logger log.Logger, etcdServers []string, prefix string) (*kitetcd.Instancer, error) {
	if instancer == nil {
		client, err := kitetcd.NewClient(ctx, etcdServers, kitetcd.ClientOptions{
			DialTimeout:             2 * time.Second,
			DialKeepAlive:           2 * time.Second,
			HeaderTimeoutPerRequest: 2 * time.Second,
		})
		if err != nil {
			logger.Log("etcd client error", err)
		}
		instancer, err = kitetcd.NewInstancer(client, prefix, logger)
		return instancer, err
	}
	return instancer, nil
}

// GetEndpoints GetEndpoints
func GetEndpoints(ctx context.Context, logger log.Logger, etcdServers []string, prefix string) ([]endpoint.Endpoint, error) {
	ins, err := getInstancer(ctx, logger, etcdServers, prefix)
	endpointer := sd.NewEndpointer(ins, func(string) (endpoint.Endpoint, io.Closer, error) {
		return endpoint.Nop, nil, nil
	}, logger)
	endpoints, err := endpointer.Endpoints()

	return endpoints, err
}
