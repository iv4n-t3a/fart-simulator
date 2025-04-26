package position_spawner

import (
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type BoundedGenerator struct {
	bounds []float64
}

func NewBoundedGenerator(bounds []float64) *BoundedGenerator {
	return &BoundedGenerator{bounds}
}

func (g *BoundedGenerator) NextPosition() vector.Vector {
	return math_util.RandVectorByVectorBounaries(vector.ZeroVector(len(g.bounds)), vector.NewVector(g.bounds))
}
