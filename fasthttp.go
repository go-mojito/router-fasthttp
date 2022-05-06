package fasthttp

import "github.com/go-mojito/mojito"

// AsDefault registers this router as the default router
func AsDefault() {
	err := mojito.Register(func() mojito.Router {
		return NewRouter()
	}, true)
	if err != nil {
		panic(err)
	}
}

// As registers this router under a given name
func As(name string) {
	err := mojito.RegisterNamed(name, func() mojito.Router {
		return NewRouter()
	}, true)
	if err != nil {
		panic(err)
	}
}
