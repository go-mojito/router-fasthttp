package fasthttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/go-mojito/mojito"
)

const (
	waitTime        = 1 * time.Second
	expectedNoError = "Expected no error, got '%s'"
	expectedBody    = "Expected body 'OK', got '%s'"
	helloWorld      = "Hello World"
)

func startServerTest(t *testing.T, r *FastHttpRouter) {
	go func() {
		if err := r.ListenAndServe(":8080"); err != nil {
			t.Errorf(expectedNoError, err)
		}
	}()
	time.Sleep(waitTime)
}

func startServerBench(b *testing.B, r *FastHttpRouter) {
	go func() {
		if err := r.ListenAndServe(":8080"); err != nil {
			b.Errorf(expectedNoError, err)
		}
	}()
	time.Sleep(waitTime)
}

func request(method string, path string) (*http.Response, error) {
	req := &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
			Path:   path,
		},
	}

	client := &http.Client{}
	return client.Do(req)
}

func Test_Router_GET(t *testing.T) {
	r := NewFastHttpRouter()
	r.GET("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("GET", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_HEAD(t *testing.T) {
	r := NewFastHttpRouter()
	r.HEAD("/", func(req mojito.Request, res mojito.Response) error {
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("HEAD", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
}

func Test_Router_POST(t *testing.T) {
	r := NewFastHttpRouter()
	r.POST("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("POST", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_PUT(t *testing.T) {
	r := NewFastHttpRouter()
	r.PUT("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("PUT", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_DELETE(t *testing.T) {
	r := NewFastHttpRouter()
	r.DELETE("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("DELETE", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_CONNECT(t *testing.T) {
	r := NewFastHttpRouter()
	r.CONNECT("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("CONNECT", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_OPTIONS(t *testing.T) {
	r := NewFastHttpRouter()
	r.OPTIONS("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("OPTIONS", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_TRACE(t *testing.T) {
	r := NewFastHttpRouter()
	r.TRACE("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("TRACE", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_PATCH(t *testing.T) {
	r := NewFastHttpRouter()
	r.PATCH("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	startServerTest(t, r)
	defer r.Shutdown()

	res, err := request("PATCH", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Test_Router_AsDefault(t *testing.T) {
	AsDefault()
	mojito.GET("/", func(req mojito.Request, res mojito.Response) error {
		res.String("OK")
		return nil
	})

	go func() {
		if err := mojito.ListenAndServe(":8080"); err != nil {
			t.Errorf(expectedNoError, err)
		}
	}()
	time.Sleep(waitTime)
	defer mojito.Shutdown()

	res, err := request("GET", "/")
	if err != nil {
		t.Errorf(expectedNoError, err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf(expectedBody, body)
	}
}

func Benchmark_Router_Handler(b *testing.B) {
	r := NewFastHttpRouter()
	r.GET("/", func(req mojito.Request, res mojito.Response) error {
		res.String(helloWorld)
		return nil
	})

	startServerBench(b, r)
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/")
	}
	b.StopTimer()
}

func Benchmark_Router_Handler_Not_Found(b *testing.B) {
	r := NewFastHttpRouter()
	r.GET("/", func(req mojito.Request, res mojito.Response) error {
		res.String(helloWorld)
		return nil
	})

	startServerBench(b, r)
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/dsfsdfa")
	}
	b.StopTimer()
}

func Benchmark_Router_Handler_With_Middleware(b *testing.B) {
	r := NewFastHttpRouter()
	r.WithMiddleware(func(ctx mojito.Context, next func() error) error {
		ctx.Metadata().Add("foo", "bar")
		return next()
	})
	r.WithMiddleware(func(ctx mojito.Context, next func() error) error {
		ctx.Metadata().Add("foo", "bar")
		return next()
	})
	r.WithMiddleware(func(ctx mojito.Context, next func() error) error {
		ctx.Metadata().Add("foo", "bar")
		return next()
	})
	r.WithMiddleware(func(ctx mojito.Context, next func() error) error {
		ctx.Metadata().Add("foo", "bar")
		return next()
	})
	r.GET("/", func(ctx mojito.Context) error {
		ctx.String(helloWorld)
		return nil
	})

	startServerBench(b, r)
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/dsfsdfa")
	}
	b.StopTimer()
}
