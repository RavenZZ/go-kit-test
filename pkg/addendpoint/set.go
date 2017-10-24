package addendpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/ravenzz/go-kit-test/pkg/service"

	"context"
	"github.com/go-kit/kit/metrics"
)

type Set struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndPoint     endpoint.Endpoint
}

func New(svc service.StringService, logger log.Logger, duration metrics.Histogram) Set {
	var uppercaseEndpoint endpoint.Endpoint
	{
		uppercaseEndpoint = MakeUppercaseEndpoint(svc)
		uppercaseEndpoint = LoggingMiddleware(log.With(logger, "method", "uppercase"))(uppercaseEndpoint)
		uppercaseEndpoint = InstrumentingMiddleware(duration.With("method", "uppercase"))(uppercaseEndpoint)
	}
	var countEndpoint endpoint.Endpoint
	{
		countEndpoint = MakeCountEndpoint(svc)
		countEndpoint = LoggingMiddleware(log.With(logger, "method", "count"))(countEndpoint)
		countEndpoint = InstrumentingMiddleware(duration.With("method", "count"))(countEndpoint)
	}

	return Set{
		UppercaseEndpoint: uppercaseEndpoint,
		CountEndPoint:     countEndpoint,
	}
}

//Uppercase(ctx context.Context, a string) (str string, err error)
//Count(ctx context.Context, a string) (count int, err error)
func (s Set) Uppercase(ctx context.Context, a string) (str string, err error) {
	resp, err := s.UppercaseEndpoint(ctx, UppercaseRequest{
		S: a,
	})
	if err != nil {
		return "", err
	}
	response := resp.(UppercaseResponse)
	return response.V, str2err(response.Err)
}

func (s Set) Count(ctx context.Context, a string) (count int, err error) {
	resp, err := s.CountEndPoint(ctx, CountRequest{
		S: a,
	})
	if err != nil {
		return 0, err
	}
	response := resp.(CountResponse)
	return response.V, str2err(response.Err)
}
