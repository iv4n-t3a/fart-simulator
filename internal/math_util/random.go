package math_util

import "math/rand"

func RandFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
