package http

import (
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	jujuratelimit "github.com/juju/ratelimit"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/ravenzz/go-kit-test/pkg/addendpoint"
	"github.com/ravenzz/go-kit-test/pkg/service"
	"github.com/sony/gobreaker"
)

func NewHTTPClient(instance string, tracer stdopentracing.Tracer, logger log.Logger) (service.StringService, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))

	var uppercaseEndpoint endpoint.Endpoint
	{
		uppercaseEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/uppercase"),
			encodeHTTPGenericRequest,
			decodeHTTPUppercaseResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(tracer, logger)),
		).Endpoint()
		uppercaseEndpoint = opentracing.TraceClient(tracer, "upper-http")(uppercaseEndpoint)
		uppercaseEndpoint = limiter(uppercaseEndpoint)
		uppercaseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "upper-http",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(uppercaseEndpoint)
	}

	var countEndpoint endpoint.Endpoint
	{
		countEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/count"),
			encodeHTTPGenericRequest,
			decodeHTTPCountResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(tracer, logger)),
		).Endpoint()
		countEndpoint = opentracing.TraceClient(tracer, "count-http")(countEndpoint)
		countEndpoint = limiter(countEndpoint)
		countEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "count-http",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(countEndpoint)
	}

	var lowercaseEndpoint endpoint.Endpoint
	{
		lowercaseEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/lower"),
			encodeHTTPGenericRequest,
			decodeHTTPLowercaseResponse,
			httptransport.ClientBefore(opentracing.ContextToHTTP(tracer, logger)),
		).Endpoint()
		lowercaseEndpoint = opentracing.TraceClient(tracer, "lower-http")(lowercaseEndpoint)
		lowercaseEndpoint = limiter(lowercaseEndpoint)
		lowercaseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "lower-http",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(lowercaseEndpoint)
	}

	return addendpoint.Set{
		UppercaseEndpoint: uppercaseEndpoint,
		CountEndPoint:     countEndpoint,
		LowercaseEndpoint: lowercaseEndpoint,
	}, nil

}
