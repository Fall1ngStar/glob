package routing

import (
	"regexp"
	"net/http"
	"html/template"
	"log"
)

func NewRouter() *RegexpRouter {
	router := RegexpRouter{}
	router.Handle("/home", HomeRoute)
	return &router
}

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

//LoadTemplates loads templates from the templates folder and automatically loads the index template
func LoadTemplates(paths ...string) *template.Template {
	t, err := template.ParseFiles("templates/index.html")
	for _, path := range paths {
		t, err = t.ParseFiles("templates/" + path)
		if err != nil {
			log.Println(err)
		}
	}
	return t
}