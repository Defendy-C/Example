package controller

import (
	"MediaWeb/api/defs"
	"net/http"
	"strings"
)

type MyRouter struct {
	routes map[string]http.HandlerFunc
}

func NewRouter()*MyRouter {
	routes := make(map[string]http.HandlerFunc, 10)
	return &MyRouter{routes:routes}
}

func(m *MyRouter)AddHandler(path string, handler http.HandlerFunc) {
	m.routes[path] = handler
}

func(m *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := strings.Trim(r.URL.Path, "/")
	paths := strings.Split(url, "/")

	if v, ok := m.routes[paths[0]];len(paths) >= 1 && ok {
		v.ServeHTTP(w, r)
	} else {
		ErrResponse(&w, defs.ErrNotPage)
	}

}