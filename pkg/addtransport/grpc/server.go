package grpc

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/ravenzz/go-kit-test/pb"
	"github.com/ravenzz/go-kit-test/pkg/addendpoint"
)

type grpcServer struct {
	uppercase grpctransport.Handler
	count     grpctransport.Handler
	lowercase grpctransport.Handler
}

func NewGRPCServer(endpoints addendpoint.Set, tracer stdopentracing.Tracer, logger log.Logger) pb.AddServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		uppercase: grpctransport.NewServer(
			endpoints.UppercaseEndpoint,
			decodeGRPCUppercaseRequest,
			encodeGRPCUppercaseResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "Upper", logger)))...,
		),
		count: grpctransport.NewServer(
			endpoints.CountEndPoint,
			decodeGRPCCountRequest,
			encodeGRPCCountResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "Count", logger)))...,
		),
	}
}

func (s *grpcServer) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseReply, error) {
	_, rep, err := s.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UppercaseReply), nil
}

func (s *grpcServer) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountReply, error) {
	_, rep, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CountReply), nil
}

func (s *grpcServer) Lowercase(ctx context.Context,req *pb.LowercaseRequest) (*pb.LowercaseResponse,error){
	_,rep, err:=s.lowercase.ServeGRPC(ctx,req)
	if err !=nil{
		return nil,err
	}
	return rep.(*pb.LowercaseResponse),nil
}


