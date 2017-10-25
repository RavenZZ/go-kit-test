package addservice

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/ravenzz/go-kit-test/pkg/service"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.StringService
}

func (mw loggingMiddleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(ctx, s)
	return
}

func (mw loggingMiddleware) Count(ctx context.Context, s string) (n int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n, err = mw.next.Count(ctx, s)
	return
}

func (mw loggingMiddleware) Lowercase(ctx context.Context,s string) (output string, err error){
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "lowercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output,err = mw.next.Lowercase(ctx,s)
	return
}
