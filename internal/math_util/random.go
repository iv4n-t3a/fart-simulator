package math_util

import (
	"fmt"
	"math/rand"

	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

func RandFloat(l float64, r float64) float64 {
	return l + rand.Float64()*(r-l)
}

func RandVectorByNumericBounaries(l float64, r float64, dim int) vector.Vector {
  vecData := make([]float64, dim)

  for i := range dim {
    vecData[i] = RandFloat(l, r)
  }

  return vector.NewVector(vecData)
}

func RandVectorByVectorBounaries(l vector.Vector, r vector.Vector) vector.Vector {
  if l.Dimensions() != r.Dimensions() {
    panic(fmt.Sprintf("Dimensions mismatch: %d != %d", l.Dimensions(), r.Dimensions()))
  }

  vecData := make([]float64, l.Dimensions())

  for i := range vecData {
    vecData[i] = RandFloat(l.Dimension(i), r.Dimension(i))
  }

  return vector.NewVector(vecData)
}
