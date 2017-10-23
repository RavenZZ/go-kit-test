package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
)

func main() {
	var (
		etcdServers = flag.String("etcd.addr", "http://10.10.100.188:2379", "etcd servers")
		prefix      = flag.String("prefix", "/permission/", "the prefix of this service ")
	)
	ctx := context.Background()
	logger := newLogger()

	etcds := strings.Split(*etcdServers, ",")

	endpoints, err := GetEndpoints(ctx, logger, etcds, *prefix)
	fmt.Println("endpoints len", len(endpoints))
	logger.Log("err", err)

}

func newLogger() log.Logger {
	return log.NewLogfmtLogger(os.Stderr)
}
