package fasthttp

import (
	"bufio"
	"net"
	"net/http"

	"github.com/valyala/fasthttp"
)

type Response struct {
	ctx       *fasthttp.RequestCtx
	header    http.Header
	headerSet bool
}

func (r *Response) Header() http.Header {
	return r.header
}

func (r *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.ctx.Conn(), bufio.NewReadWriter(bufio.NewReader(r.ctx.Conn()), bufio.NewWriter(r.ctx.Conn())), nil
}

func (r *Response) Write(body []byte) (int, error) {
	if !r.headerSet {
		r.setHeaders(body)
		r.headerSet = true
	}
	return r.ctx.Write(body)
}

func (r *Response) WriteHeader(status int) {
	r.ctx.SetStatusCode(status)
}

func (r *Response) setHeaders(body []byte) {
	// Because of special header handling in fasthttp, we need to manually
	// set the header here.
	for k, vv := range r.header {
		for _, v := range vv {
			r.ctx.Response.Header.Add(k, v)
		}
	}
}

func NewResponse(ctx *fasthttp.RequestCtx) *Response {
	return &Response{
		ctx:    ctx,
		header: make(http.Header),
	}
}
