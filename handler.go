package fasthttp

import (
	"fmt"
	"net/http"

	"github.com/go-mojito/mojito/log"
	"github.com/go-mojito/mojito/pkg/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// WithMojitoHandler converts a Mojito handler to a fasthttp handler
func WithMojitoHandler(handler router.Handler) fasthttp.RequestHandler {
	return func(reqCtx *fasthttp.RequestCtx) {
		httpReq := &http.Request{}
		if err := fasthttpadaptor.ConvertRequest(reqCtx, httpReq, true); err != nil {
			reqCtx.Error(err.Error(), http.StatusInternalServerError)
			return
		}
		req := router.NewRequest(httpReq)

		params := make(map[string]string)
		reqCtx.VisitUserValues(func(key []byte, value interface{}) {
			params[string(key)] = fmt.Sprint(value)
		})
		req.SetParams(params)

		res := router.NewResponse(NewResponse(reqCtx))
		ctx := router.NewContext(req, res)
		if err := handler.Serve(ctx); err != nil {
			log.Field("router", "fasthttp").Error(err)
		}
	}
}
