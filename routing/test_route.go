package routing

import (
	"net/http"
	"html/template"
	"fmt"
	"glob/models"
)

type Person struct {
	Name string
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/content.html")
	person := Person{r.URL.Path[len("/test/"):]}
	err = t.Execute(w, models.PageData{Title: "Hello there", Data: person})
	if err != nil {
		fmt.Println(err)
	}
}
