package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ravenzz/go-kit-test/pkg/addendpoint"
	"github.com/ravenzz/go-kit-test/pkg/addservice"
)

// Error Start

type errorWrapper struct {
	Error string `json:"error"`
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case addservice.ErrTwoZeroes, addservice.ErrMaxSizeExceeded, addservice.ErrIntOverflow:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// Error End

// Encode Genric Start

// encodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(addendpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// Encode Genric End

// Decode Start

func decodeHTTPUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addendpoint.UppercaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPUppercaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp addendpoint.UppercaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeHTTPCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addendpoint.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPCountResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp addendpoint.CountRequest
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeHTTPLowercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addendpoint.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPLowercaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp addendpoint.LowercaseRequest
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// Decode End
