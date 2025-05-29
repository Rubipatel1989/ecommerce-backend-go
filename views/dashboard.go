package views

import (
	"html/template"
	"log"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/header.html",
		"templates/menu.html",
		"templates/footer.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not load dashboard", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "header.html", nil)
}
