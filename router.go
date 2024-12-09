package fasthttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	fhttpRouter "github.com/fasthttp/router"
	"github.com/go-mojito/mojito/log"
	"github.com/go-mojito/mojito/pkg/router"
	"github.com/infinytum/structures"
	"github.com/valyala/fasthttp"
	fhttp "github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Router is a router implementation based on fasthttp
type Router struct {
	Middleware []interface{}
	Routes     structures.Table[string, string, router.Handler]
	*fhttp.Server
	*fhttpRouter.Router

	ErrorHandler router.Handler
}

// NewRouter will create new instance of the mojito fasthttp router implementation
func NewRouter() *Router {
	return &Router{
		Middleware: make([]interface{}, 0),
		Routes:     structures.NewTable[string, string, router.Handler](),
		Router:     fhttpRouter.New(),
	}
}

// GET will add a route to this router for the get method
func (r *Router) GET(path string, handler interface{}) error {
	return r.WithRoute(http.MethodGet, path, handler)
}

// HEAD will add a route to this router for the head method
func (r *Router) HEAD(path string, handler interface{}) error {
	return r.WithRoute(http.MethodHead, path, handler)
}

// POST will add a route to this router for the post method
func (r *Router) POST(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPost, path, handler)
}

// PUT will add a route to this router for the put method
func (r *Router) PUT(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPut, path, handler)
}

// DELETE will add a route to this router for the delete method
func (r *Router) DELETE(path string, handler interface{}) error {
	return r.WithRoute(http.MethodDelete, path, handler)
}

// CONNECT will add a route to this router for the connect method
func (r *Router) CONNECT(path string, handler interface{}) error {
	return r.WithRoute(http.MethodConnect, path, handler)
}

// OPTIONS will add a route to this router for the options method
func (r *Router) OPTIONS(path string, handler interface{}) error {
	return r.WithRoute(http.MethodOptions, path, handler)
}

// TRACE will add a route to this router for the trace method
func (r *Router) TRACE(path string, handler interface{}) error {
	return r.WithRoute(http.MethodTrace, path, handler)
}

// PATCH will add a route to this router for the patch method
func (r *Router) PATCH(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPatch, path, handler)
}

// WithGroup will create a new route group for the given prefix
func (r *Router) WithGroup(path string, callback func(group router.Group)) error {
	rg := router.NewGroup()
	callback(rg)
	return rg.ApplyToRouter(r, path)
}

// WithRoute will add a new route with the given RouteMethod to the router
func (r *Router) WithRoute(method string, path string, handler interface{}) error {
	h, err := router.GetOrCreateHandler(handler)
	if err != nil {
		log.With("method", method, "path", path, "err", err).Error("Error creating route handler")
		return err
	}

	// Chain router-wide middleware to the (new) handler
	for _, middleware := range r.Middleware {
		if err := h.AddMiddleware(middleware); err != nil {
			log.With("method", method, "path", path, "err", err).Error("Failed to chain middleware to route")
			return err
		}
	}

	if regexp, err := regexp.Compile(`:(\w{1,})`); err == nil {
		path = regexp.ReplaceAllString(path, "{$1}")
	}
	if regexp, err := regexp.Compile(`\*(\w{1,})`); err == nil {
		path = regexp.ReplaceAllString(path, "{$1:*}")
	}

	switch method {
	case http.MethodGet:
		r.Router.GET(path, r.withMojitoHandler(h))
	case http.MethodHead:
		r.Router.HEAD(path, r.withMojitoHandler(h))
	case http.MethodPost:
		r.Router.POST(path, r.withMojitoHandler(h))
	case http.MethodPut:
		r.Router.PUT(path, r.withMojitoHandler(h))
	case http.MethodDelete:
		r.Router.DELETE(path, r.withMojitoHandler(h))
	case http.MethodConnect:
		r.Router.CONNECT(path, r.withMojitoHandler(h))
	case http.MethodOptions:
		r.Router.OPTIONS(path, r.withMojitoHandler(h))
	case http.MethodTrace:
		r.Router.TRACE(path, r.withMojitoHandler(h))
	case http.MethodPatch:
		r.Router.PATCH(path, r.withMojitoHandler(h))
	default:
		log.With("method", method, "path", path).Error("The fasthttp router implementation unfortunately does not support this HTTP method")
		return errors.New("the given HTTP method is not available on this router")
	}
	return r.Routes.Add(path, method, h)
}

// WithNotFoundHandler will set the not found handler for the router
func (r *Router) WithNotFoundHandler(handler interface{}) error {
	if h, err := router.GetOrCreateHandler(handler); err != nil {
		log.With("err", err).Error("Error creating not found handler")
		return err
	} else {
		r.Router.NotFound = r.withMojitoHandler(h)
	}
	return nil
}

// WithMethodNotAllowedHandler will set the not allowed handler for the router
func (r *Router) WithMethodNotAllowedHandler(handler interface{}) error {
	if h, err := router.GetOrCreateHandler(handler); err != nil {
		log.With("err", err).Error("Error creating not allowed handler")
		return err
	} else {
		r.Router.MethodNotAllowed = r.withMojitoHandler(h)
	}
	return nil
}

// WithErrorHandler will set the error handler for the router
func (r *Router) WithErrorHandler(handler interface{}) error {
	if h, err := router.GetOrCreateHandler(handler); err != nil {
		log.With("err", err).Error("Error creating error handler")
		return err
	} else {
		r.ErrorHandler = h
		r.Router.PanicHandler = r.onError
	}
	return nil
}

// WithMiddleware will add a middleware to the router
func (r *Router) WithMiddleware(handler interface{}) error {
	for _, rm := range r.Routes.ToMap() {
		for _, h := range rm {
			if err := h.AddMiddleware(handler); err != nil {
				log.With("err", err).Error("Failed to chain middleware to route")
				return err
			}
		}
	}
	r.Middleware = append(r.Middleware, handler)
	return nil
}

// ListenAndServe will start an HTTP webserver on the given address
func (r *Router) ListenAndServe(address string) error {
	r.Server = &fhttp.Server{
		Handler: r.Router.Handler,
	}
	return r.Server.ListenAndServe(address)
}

// Shutdown will gracefully shutdown the router
func (r *Router) Shutdown() error {
	return r.Server.Shutdown()
}

func (r *Router) onError(reqCtx *fhttp.RequestCtx, rec interface{}) {
	httpReq := &http.Request{}
	if err := fasthttpadaptor.ConvertRequest(reqCtx, httpReq, true); err != nil {
		reqCtx.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	mreq := router.NewRequest(httpReq)
	mres := router.NewResponse(NewResponse(reqCtx))
	ctx, cancel := router.NewContext(mreq, mres)
	cancel(fmt.Errorf("%v", rec))
	if err := r.ErrorHandler.Serve(ctx); err != nil {
		panic(err)
	}
}

func (r *Router) withMojitoHandler(handler router.Handler) fasthttp.RequestHandler {
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
		ctx, cancel := router.NewContext(req, res)
		cancel(handler.Serve(ctx))
		if ctx.Err() != context.Canceled {
			if r.ErrorHandler != nil {
				if err := r.ErrorHandler.Serve(ctx); err != context.Canceled && err != nil {
					panic(err)
				}
			} else {
				panic(ctx.Err())
			}
		}
	}
}
