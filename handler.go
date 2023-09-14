package handler

import (
	"context"
	"net/http"
)
type Response struct {
	Body []byte
	StatusCode int
	Header http.Header
}

type Request struct {
	Body        []byte
	Header      http.Header
	QueryString string
	Method      string
	Host        string
        Path        string
	ctx         context.Context
}

func (r *Request) Context() context.Context {
	return r.ctx
}

func (r *Request) WithContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	r.ctx = ctx
}

type FunctionHandler interface {
	Handle(req Request) (Response, error)
}

func init() {

}

