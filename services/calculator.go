package services

import "EduScoreCalc/models"

func CalculateWeightedAverage(grades []models.Grade, weights map[string]int) float64 {
	totalScore := 0
	totalWeight := 0

	for _, grade := range grades {
		weight := weights[grade.ActivityCode]
		totalScore += grade.Score * weight
		totalWeight += weight
	}

	if totalWeight == 0 {
		return 0
	}
	return float64(totalScore) / float64(totalWeight)
}
