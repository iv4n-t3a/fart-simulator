package math_util

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/config"
)

// Solves ax^2 + bx + c = 0
func SolveSqrEquivaltion(a float64, b float64, c float64) (*float64, *float64) {
	d := Discriminant(a, b, c)

	if math.Abs(d) <= config.Eps {
		x := -(b + math.Sqrt(d)) / 2 / a
		return &x, nil
	}
	if d < 0 {
		return nil, nil
	}

	x1 := -(b + math.Sqrt(d)) / 2 / a
	x2 := -(b - math.Sqrt(d)) / 2 / a
	return &x1, &x2
}

func Discriminant(a float64, b float64, c float64) float64 {
	return b*b - 4*a*c
}
