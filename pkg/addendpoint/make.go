package addendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ravenzz/go-kit-test/pkg/service"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}

}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v, err := svc.Count(ctx, req.S)
		if err != nil {
			return countResponse{v, err.Error()}, nil
		}
		return countResponse{v, ""}, nil
	}
}
