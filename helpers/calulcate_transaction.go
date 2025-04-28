package helpers

import (
	"fmt"
	"math"
)

func PatternCheck(baselline float64, amount float64) (float64, float64) {
	percentage := (amount - baselline) / baselline * 100
	score := math.Min(100, percentage/4.0)
	return percentage, score
}

func FrequencyCheck(count int64) (string, string) {
	var score string
	narration := fmt.Sprintf("Order frequency: %s order in 1 hour", count)
	if count > 8 {
		score = "90-100"
	} else if count >= 7 && count <= 8 {
		score = "80-89"
	} else if count >= 6 {
		score = "70-79"
	} else if count >= 5 {
		score = "50-69"
	} else {
		score = "50"
	}

	return score, narration
}
