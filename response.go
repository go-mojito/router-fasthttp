package fasthttp

import (
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
	hasContentType := false
	for k, vv := range r.header {
		if k == fasthttp.HeaderContentType {
			hasContentType = true
		}

		for _, v := range vv {
			r.ctx.Response.Header.Add(k, v)
		}
	}
	if !hasContentType {
		// From net/http.ResponseWriter.Write:
		// If the Header does not contain a Content-Type line, Write adds a Content-Type set
		// to the result of passing the initial 512 bytes of written data to DetectContentType.
		l := 512
		if len(body) < 512 {
			l = len(body)
		}
		r.ctx.Response.Header.Set(fasthttp.HeaderContentType, http.DetectContentType(body[:l]))
	}
}

func NewResponse(ctx *fasthttp.RequestCtx) *Response {
	return &Response{
		ctx:    ctx,
		header: make(http.Header),
	}
}
