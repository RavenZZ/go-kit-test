package grpc

import (
	"context"

	"github.com/ravenzz/go-kit-test/pb"

	"github.com/ravenzz/go-kit-test/pkg/addendpoint"
)

// decodeGRPC Request
// encodeGRPC Request
// decodeGRPC Response
// encodeGRPC Response

// Uppercase  Start

func decodeGRPCUppercaseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UppercaseRequest)
	return addendpoint.UppercaseRequest{S: req.A}, nil
}

func encodeGRPCUppercaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(addendpoint.UppercaseRequest)
	return &pb.UppercaseRequest{A: req.S}, nil
}

func decodeGRPCUppercaseResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	resp := grpcReply.(*pb.UppercaseReply)
	return addendpoint.UppercaseResponse{
		V:   resp.Str,
		Err: resp.Err,
	}, nil
}

func encodeGRPCUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.UppercaseResponse)
	return &pb.UppercaseReply{
		Str: resp.V,
		Err: resp.Err,
	}, nil
}

// Uppercase  End

// Count Start

func decodeGRPCCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CountRequest)
	return addendpoint.CountRequest{S: req.A}, nil
}

func encodeGRPCCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(addendpoint.CountRequest)
	return &pb.CountRequest{A: req.S}, nil
}

func decodeGRPCCountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	resp := grpcReply.(*pb.CountReply)
	return addendpoint.CountResponse{
		V:   int(resp.Count),
		Err: resp.Err,
	}, nil
}

func encodeGRPCCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.CountResponse)
	return &pb.CountReply{
		Count: int64(resp.V),
		Err:   resp.Err,
	}, nil
}

// Count End


// Uppercase  Start

func encodeGRPCLowercaseRequest(_ context.Context, grpcReq interface{}) (interface{},error){
	req:= grpcReq.(*pb.LowercaseRequest)
	return addendpoint.LowercaseRequest{S:req.A},nil
}

func decodeGRPCLowercaseRequest(_ context.Context, request interface{}) (interface{},error) {
	req := request.(addendpoint.LowercaseRequest)
	return &pb.LowercaseRequest{A: req.S}, nil
}

func encodeGRPCLowercaseResponse(_ context.Context, grpcReply interface{}) (interface{},error){
	resp:= grpcReply.(*pb.LowercaseResponse)
	return addendpoint.LowercaseResponse{V:resp.Str,Err:resp.Err},nil
}

func decodeGRPCLowercaseResponse(_ context.Context, response interface{}) (interface{},error){
	resp:= response.(addendpoint.LowercaseResponse)
	return &pb.LowercaseResponse{
		Str:resp.V,
		Err:resp.Err,
	},nil
}

// Lowercase  End


