package velocity_spawner

import (
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type NaiveVelocitySpawner struct {
	maxVelocity float64
	dim         int
}

func NewNaiveVelocitySpawner(maxVelocity float64, dim int) VelocitySpawner {
	return &NaiveVelocitySpawner{
		maxVelocity: maxVelocity,
		dim:         dim,
	}
}

func (s *NaiveVelocitySpawner) NextVelocity() vector.Vector {
	return math_util.RandVectorByNumericBounaries(-s.maxVelocity, s.maxVelocity, s.dim)
}
