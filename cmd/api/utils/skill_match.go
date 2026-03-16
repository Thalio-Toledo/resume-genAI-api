package utils

import "math"

func CosineSimilarity(a, b []float32) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if len(a) != len(b) {
		return 0
	}

	var dot, normA, normB float64

	for i := 0; i < len(a); i++ {
		dot += float64(a[i] * b[i])
		normA += float64(a[i] * a[i])
		normB += float64(b[i] * b[i])
	}

	denominator := math.Sqrt(normA) * math.Sqrt(normB)
	if denominator == 0 {
		return 0
	}

	return dot / denominator
}

func MatchEmbedding(a, b []float32, threshold float64) bool {
	score := CosineSimilarity(a, b)
	return score >= threshold
}
