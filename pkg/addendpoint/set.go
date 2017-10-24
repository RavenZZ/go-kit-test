package addendpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/ravenzz/go-kit-test/pkg/service"

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
