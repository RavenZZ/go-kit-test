package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-kit/kit/log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// go run *.go -instance=127.0.0.1:40001 -port=:40001 -accessAddr=http://127.0.0.1:40001
// go run *.go -instance=127.0.0.1:40002 -port=:40002 -accessAddr=http://127.0.0.1:40002
// go run *.go -instance=127.0.0.1:40003 -port=:40003 -accessAddr=http://127.0.0.1:40003

func main() {

	var (
		etcdServers = flag.String("etcd.addr", "http://10.10.100.188:2379", "etcd servers")
		prefix      = flag.String("prefix", "/permission/", "the prefix of this service ")
		instance    = flag.String("instance", "127.0.0.1:40001", "")
		port        = flag.String("port", ":40001", "port to listen")
		accessAddr  = flag.String("accessAddr", "http://127.0.0.1:40001", "")
	)
	flag.Parse()
	errChan := make(chan error)

	ctx := context.Background()
	// 定义logger Start
	logger := newLogger()
	// 定义logger End

	// 定义Metrics Start
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here
	// 定义Metrics End

	var svc StringService
	svc = stringService{}

	// 注入logging
	svc = loggingMiddleware{
		logger: logger,
		next:   svc,
	}

	// 注入metrics
	svc = instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           svc,
	}

	etcds := strings.Split(*etcdServers, ",")

	// 声明 Register
	registar := Register(ctx, logger, etcds, *prefix, *instance, *accessAddr)

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	// metrics 查询
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		registar.Register()
		logger.Log("MSG", fmt.Sprintf("Starting server at port:%v", *port))
		errChan <- http.ListenAndServe(*port, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			syscall.SIGINT,
			syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err := <-errChan
	registar.Deregister()
	logger.Log("Error", err)
	os.Exit(0)
}

func newLogger() log.Logger {
	return log.NewLogfmtLogger(os.Stderr)
}
