package hayane

import (
	"net/http"

	"github.com/nazhard/chiyo"
)

type Haya struct {
	router *chiyo.Router
}

func New() *Haya {
	r := chiyo.NewRouter()
	return &Haya{router: r}
}

func (h *Haya) GET(path string, handler func(*Context)) {
	h.router.AddRoute(http.MethodGet, path, func(w http.ResponseWriter, req *http.Request) {
		c := &Context{Writer: w, Request: req}
		handler(c)
	})
}

func (h *Haya) POST(path string, handler func(c *Context)) {
	h.router.AddRoute(http.MethodPost, path, func(w http.ResponseWriter, req *http.Request) {
		c := &Context{Writer: w, Request: req}
		handler(c)
	})
}

func (h *Haya) PUT(path string, handler func(*Context)) {
	h.router.AddRoute(http.MethodPut, path, func(w http.ResponseWriter, req *http.Request) {
		c := &Context{Writer: w, Request: req}
		handler(c)
	})
}

func (h *Haya) DELETE(path string, handler func(*Context)) {
	h.router.AddRoute(http.MethodDelete, path, func(w http.ResponseWriter, req *http.Request) {
		c := &Context{Writer: w, Request: req}
		handler(c)
	})
}

func (h *Haya) Use(middleware func(http.HandlerFunc) http.HandlerFunc) {
	h.router.Use(middleware)
}

func (h *Haya) Start(address string) error {
	return http.ListenAndServe(address, h.router)
}
