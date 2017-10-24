package addservice

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/ravenzz/go-kit-test/pkg/service"
)

// New new
func New(logger log.Logger, requestCount metrics.Counter, requestLatency metrics.Histogram) service.StringService {
	var svc service.StringService
	// 注入logging
	svc = loggingMiddleware{
		logger: logger,
		next:   svc,
	}

	// 注入metrics
	svc = instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           svc,
	}

	return svc
}
