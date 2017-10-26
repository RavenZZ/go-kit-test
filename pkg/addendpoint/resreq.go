package addendpoint

import "errors"

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type LowercaseRequest struct {
	S string `json:"s"`
}

type LowercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so if they've
// failed, and if so encode them using a separate write path based on the error.
type Failer interface {
	Failed() error
}

func (r UppercaseResponse) Failed() error {
	return errors.New(r.Err)
}

func (r CountResponse) Failed() error {
	return errors.New(r.Err)
}

func (r LowercaseResponse) Failed() error {
	return errors.New(r.Err)
}
