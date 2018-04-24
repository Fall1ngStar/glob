package routing

import (
	"net/http"
	"glob/models"
	"log"
)

type Person struct {
	Name string
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	t := LoadTemplates("content.html")
	person := Person{r.URL.Path[len("/test/"):]}
	err := t.Execute(w, models.PageData{Title: "Hello there", Data: person})
	if err != nil {
		log.Println(err)
	}
}
