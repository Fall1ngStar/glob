package routing

import (
	"regexp"
	"net/http"
)

type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpRouter struct {
	routes []*Route
}

func (h *RegexpRouter) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &Route{regexp.MustCompile(pattern), http.HandlerFunc(handler)})
}

func (h *RegexpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func NewRouter() *RegexpRouter {
	router := RegexpRouter{}
	router.Handle("/test/*", TestRoute)
	return &router
}
