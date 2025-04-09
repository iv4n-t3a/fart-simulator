package mathutil_test

import (
	"math"
	"testing"

	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/stretchr/testify/assert"
)

func TestSimpleEquation(t *testing.T) {
  x1, x2, ok := math_util.SolveSqrEquation(1.0, -4.0, 3.0)
  assert.True(t, ok, "Solutions not found")
  assert.InEpsilon(t, math.Min(x1, x2), 1.0, 0.001, "Min solution doesn't match")
  assert.InEpsilon(t, math.Max(x1, x2), 3.0, 0.001, "Max solution doesn't match")
}

func TestOnlyOneSolution(t *testing.T) {
  x1, x2, ok := math_util.SolveSqrEquation(1.0, -4.0, 4.0)
  assert.True(t, ok, "Solutions not found")
  assert.InEpsilon(t, x1, 2.0, 0.001, "Solution doesn't match")
  assert.InEpsilon(t, x2, 2.0, 0.001, "Solution doesn't match")
}

func TestNoSolutions(t *testing.T) {
  _, _, ok := math_util.SolveSqrEquation(1.0, -4.0, 5.0)
  assert.False(t, ok, "Extra solutions found")
}
