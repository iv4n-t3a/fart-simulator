package mathutil_test

import (
	"math"
	"testing"

	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/stretchr/testify/assert"
)

func TestSimpleEqualation(t *testing.T) {
  x1, x2 := math_util.SolveSqrEquivaltion(1.0, -4.0, 3.0)
  assert.InEpsilon(t, math.Min(*x1, *x2), 1.0, 0.001, "Min solution doesn't match")
  assert.InEpsilon(t, math.Max(*x1, *x2), 3.0, 0.001, "Max solution doesn't match")
}

func TestOnlyOneSolution(t *testing.T) {
  x1, x2 := math_util.SolveSqrEquivaltion(1.0, -4.0, 4.0)
  assert.InEpsilon(t, *x1, 2.0, 0.001, "Solution doesn't match")
  print(x2)
  assert.Nil(t, x2, "Extra solution found")
}

func TestNoSolutions(t *testing.T) {
  x1, x2 := math_util.SolveSqrEquivaltion(1.0, -4.0, 5.0)
  assert.Nil(t, x1, "Extra solution found")
  assert.Nil(t, x2, "Extra solution found")
}
