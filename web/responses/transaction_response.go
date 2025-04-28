package responses

type FrequencyCheck struct {
	Score     string `json:"score"`
	Narration string `json:"narration"`
}

type PatternCheck struct {
	Score      float64 `json:"score"`
	Percentage float64 `json:"percentage"`
}

type TransactionResponse struct {
	TransactionId   int     `json:"transaction_id"`
	FraudScore      float64 `json:"fraud_score"`
	RiskLevel       string  `json:"risk_level"`
	DetectionResult struct {
		FrequencyCheck FrequencyCheck `json:"frequency_check"`
		PatternCheck   PatternCheck   `json:"pattern_check"`
	} `json:"detection_results"`
}
