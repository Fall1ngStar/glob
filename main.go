package main

import (
	"net/http"
	"log"
	"regexp"
	"html/template"
	"fmt"
	"strings"
)

type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*Route
}

func (h *RegexpHandler) Handle(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &Route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

type Person struct {
	FirstName, LastName, Url string
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handler")
	t, err := template.ParseFiles("templates/test.html")
	fmt.Println(t)
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Split(r.URL.Path, "/")[2:]
	err = t.Execute(w, Person{arr[0], arr[1], r.URL.Path})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handler := RegexpHandler{}
	handler.Handle(regexp.MustCompile("/edit/*"), editHandler)
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
