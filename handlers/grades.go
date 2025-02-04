package handlers

import (
	"html/template"
	"net/http"
)

func GradesHandler(w http.ResponseWriter, r *http.Request) {
	schoolType := r.URL.Query().Get("type")
	tmpl := template.Must(template.ParseFiles("templates/grades.html"))

	data := struct {
		SchoolType string
	}{
		SchoolType: schoolType,
	}

	tmpl.Execute(w, data)
}
