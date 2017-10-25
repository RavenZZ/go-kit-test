package addendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ravenzz/go-kit-test/pkg/service"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}

}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v, err := svc.Count(ctx, req.S)
		if err != nil {
			return CountResponse{v, err.Error()}, nil
		}
		return CountResponse{v, ""}, nil
	}
}


func MakeLowercaseEndpoint(svc service.StringService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req:= request.(LowercaseRequest)
		v,err := svc.Lowercase(ctx,req.S)
		if err !=nil{
			return  LowercaseResponse{V:v,Err:err.Error()},nil
		}
		return  LowercaseResponse{V:v,Err:""},nil
	}
}