package main

import (
	"EduScoreCalc/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/grades", handlers.GradesHandler)
	http.HandleFunc("/calculate", handlers.CalculateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
