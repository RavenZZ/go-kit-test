package http

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/ravenzz/go-kit-test/pkg/addendpoint"

	"net/http"

	"github.com/go-kit/kit/tracing/opentracing"
)

func NewHTTPHandler(endpoints addendpoint.Set, tracer stdopentracing.Tracer, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	m := http.NewServeMux()
	m.Handle("/uppercase", httptransport.NewServer(
		endpoints.UppercaseEndpoint,
		decodeHTTPUppercaseRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(tracer, "upper-http", logger)))...,
	))

	m.Handle("/count", httptransport.NewServer(
		endpoints.CountEndPoint,
		decodeHTTPCountRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(tracer, "count-http", logger)))...,
	))

	m.Handle("/lowercase", httptransport.NewServer(
		endpoints.LowercaseEndpoint,
		decodeHTTPLowercaseRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(tracer, "lower-http", logger)))...,
	))

	return m
}
