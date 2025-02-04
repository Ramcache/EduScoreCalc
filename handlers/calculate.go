package handlers

import (
	"EduScoreCalc/data"
	"EduScoreCalc/models"
	"EduScoreCalc/services"
	"html/template"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	schoolType := r.FormValue("schoolType")
	activityCodes := r.Form["activityCode"]
	scores := r.Form["score"]

	var grades []models.Grade
	weights := data.PrimaryWeights

	if schoolType == "secondary" {
		weights = data.SecondaryWeights
	}

	for i := range activityCodes {
		score, _ := strconv.Atoi(scores[i])
		grades = append(grades, models.Grade{
			ActivityCode: activityCodes[i],
			Score:        score,
		})
	}

	result := services.CalculateWeightedAverage(grades, weights)

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, struct {
		Result float64
	}{
		Result: result,
	})
}
