package fasthttp

import (
	"errors"
	"net/http"
	"regexp"

	fhttpRouter "github.com/fasthttp/router"
	"github.com/go-mojito/mojito"
	"github.com/go-mojito/mojito/log"
	"github.com/go-mojito/mojito/pkg/router"
	"github.com/infinytum/structures"
	fhttp "github.com/valyala/fasthttp"
)

// FastHttpRouter is a router implementation based on fasthttp
type FastHttpRouter struct {
	Miiddleware []interface{}
	Routes      structures.Table[string, string, router.Handler]
	*fhttp.Server
	*fhttpRouter.Router
}

// NewFastHttpRouter will create new instance of the mojito fasthttp router implementation
func NewFastHttpRouter() mojito.Router {
	return &FastHttpRouter{
		Miiddleware: make([]interface{}, 0),
		Routes:      structures.NewTable[string, string, router.Handler](),
		Router:      fhttpRouter.New(),
	}
}

// GET will add a route to this router for the get method
func (r *FastHttpRouter) GET(path string, handler interface{}) error {
	return r.WithRoute(http.MethodGet, path, handler)
}

// HEAD will add a route to this router for the head method
func (r *FastHttpRouter) HEAD(path string, handler interface{}) error {
	return r.WithRoute(http.MethodHead, path, handler)
}

// POST will add a route to this router for the post method
func (r *FastHttpRouter) POST(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPost, path, handler)
}

// PUT will add a route to this router for the put method
func (r *FastHttpRouter) PUT(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPut, path, handler)
}

// DELETE will add a route to this router for the delete method
func (r *FastHttpRouter) DELETE(path string, handler interface{}) error {
	return r.WithRoute(http.MethodDelete, path, handler)
}

// CONNECT will add a route to this router for the connect method
func (r *FastHttpRouter) CONNECT(path string, handler interface{}) error {
	return r.WithRoute(http.MethodConnect, path, handler)
}

// OPTIONS will add a route to this router for the options method
func (r *FastHttpRouter) OPTIONS(path string, handler interface{}) error {
	return r.WithRoute(http.MethodOptions, path, handler)
}

// TRACE will add a route to this router for the trace method
func (r *FastHttpRouter) TRACE(path string, handler interface{}) error {
	return r.WithRoute(http.MethodTrace, path, handler)
}

// PATCH will add a route to this router for the patch method
func (r *FastHttpRouter) PATCH(path string, handler interface{}) error {
	return r.WithRoute(http.MethodPatch, path, handler)
}

// WithDefaultHandler will set the default handler for the router
func (r *FastHttpRouter) WithDefaultHandler(handler interface{}) error {
	if h, err := router.GetOrCreateHandler(handler); err != nil {
		log.Errorf("Error creating default handler: %s", err)
		return err
	} else {
		r.Router.NotFound = WithMojitoHandler(h)
	}
	return nil
}

// WithErrorHandler will set the error handler for the router
func (r *FastHttpRouter) WithErrorHandler(handler interface{}) error {
	if h, err := router.GetOrCreateHandler(handler); err != nil {
		log.Errorf("Error creating error handler: %s", err)
		return err
	} else {
		r.Router.PanicHandler = func(ctx *fhttp.RequestCtx, _ interface{}) {
			WithMojitoHandler(h)(ctx)
		}
	}
	return nil
}

// WithGroup will create a new route group for the given prefix
func (r *FastHttpRouter) WithGroup(path string, callback func(group router.Group)) error {
	rg := router.NewGroup()
	callback(rg)
	return rg.ApplyToRouter(r, path)
}

// WithMiddleware will add a middleware to the router
func (r *FastHttpRouter) WithMiddleware(handler interface{}) error {
	for _, rm := range r.Routes.ToMap() {
		for _, h := range rm {
			if err := h.AddMiddleware(handler); err != nil {
				log.Error(err)
				return err
			}
		}
	}
	r.Miiddleware = append(r.Miiddleware, handler)
	return nil
}

// WithRoute will add a new route with the given RouteMethod to the router
func (r *FastHttpRouter) WithRoute(method string, path string, handler interface{}) error {
	h, err := router.GetOrCreateHandler(handler)
	if err != nil {
		log.Field("method", method).Field("path", path).Errorf("Error creating rotue handler: %s", err)
		return err
	}

	// Chain router-wide middleware to the (new) handler
	for _, middleware := range r.Miiddleware {
		if err := h.AddMiddleware(middleware); err != nil {
			log.Field("method", method).Field("path", path).Errorf("Failed to chain middleware to route: %s", err)
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
		r.Router.GET(path, WithMojitoHandler(h))
	case http.MethodHead:
		r.Router.HEAD(path, WithMojitoHandler(h))
	case http.MethodPost:
		r.Router.POST(path, WithMojitoHandler(h))
	case http.MethodPut:
		r.Router.PUT(path, WithMojitoHandler(h))
	case http.MethodDelete:
		r.Router.DELETE(path, WithMojitoHandler(h))
	case http.MethodConnect:
		r.Router.CONNECT(path, WithMojitoHandler(h))
	case http.MethodOptions:
		r.Router.OPTIONS(path, WithMojitoHandler(h))
	case http.MethodTrace:
		r.Router.TRACE(path, WithMojitoHandler(h))
	case http.MethodPatch:
		r.Router.PATCH(path, WithMojitoHandler(h))
	default:
		log.Field("method", method).Field("path", path).Error("The fasthttp router implementation unfortunately does not support this HTTP method")
		return errors.New("the given HTTP method is not available on this router")
	}
	r.Routes.Add(path, method, h)
	return nil
}

// ListenAndServe will start an HTTP webserver on the given address
func (r *FastHttpRouter) ListenAndServe(address string) error {
	r.Server = &fhttp.Server{
		Handler: r.Router.Handler,
	}
	return r.Server.ListenAndServe(address)
}

// Shutdown will gracefully shutdown the router
func (r *FastHttpRouter) Shutdown() error {
	return r.Server.Shutdown()
}
