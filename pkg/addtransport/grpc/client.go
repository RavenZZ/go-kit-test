package grpc

import (
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	jujuratelimit "github.com/juju/ratelimit"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/ravenzz/go-kit-test/pb"
	"github.com/ravenzz/go-kit-test/pkg/addendpoint"
	"github.com/ravenzz/go-kit-test/pkg/service"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) service.StringService {
	limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(1000, 1000))

	var upperEndpoint endpoint.Endpoint
	{
		upperEndpoint = grpctransport.NewClient(
			conn,
			"pb.StrService",
			"Upper",
			encodeGRPCUppercaseRequest,
			decodeGRPCUppercaseResponse,
			pb.UppercaseReply{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		upperEndpoint = opentracing.TraceClient(tracer, "upper-grpc")(upperEndpoint)
		upperEndpoint = limiter(upperEndpoint)
		upperEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "upper-grpc",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(upperEndpoint)
	}

	var countEndpoint endpoint.Endpoint
	{
		countEndpoint = grpctransport.NewClient(
			conn,
			"pb.StrService",
			"Count",
			encodeGRPCCountRequest,
			decodeGRPCCountResponse,
			pb.CountReply{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		countEndpoint = opentracing.TraceClient(tracer, "count-grpc")(countEndpoint)
		countEndpoint = limiter(countEndpoint)
		countEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "count-grpc",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(countEndpoint)
	}

	var lowerEndpoint endpoint.Endpoint
	{
		lowerEndpoint = grpctransport.NewClient(
			conn,
			"pb.StrService",
			"Lower",
			encodeGRPCLowercaseRequest,
			decodeGRPCLowercaseResponse,
			pb.LowercaseResponse{},
			grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
		).Endpoint()
		countEndpoint = opentracing.TraceClient(tracer, "lower-grpc")(countEndpoint)
		countEndpoint = limiter(countEndpoint)
		countEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "lower-grpc",
			MaxRequests: 5,
			Interval:    20 * time.Second,
			Timeout:     20 * time.Second,
			ReadyToTrip: nil,
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				// 这里应该提交报警, 状态异常或恢复
			},
		}))(countEndpoint)
	}

	return addendpoint.Set{
		UppercaseEndpoint: upperEndpoint,
		CountEndPoint:     countEndpoint,
		LowercaseEndpoint: lowerEndpoint,
	}

}
