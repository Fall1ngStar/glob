package routing

import (
	"net/http"
	"glob/models"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	t := LoadTemplates("home.html")
	t.Execute(w, models.PageData{
		Title: "Home",
	})
}
