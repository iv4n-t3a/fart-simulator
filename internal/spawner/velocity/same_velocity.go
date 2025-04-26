package velocity_spawner

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type SameVelocitySpawner struct {
	velocity float64
	dim      int
}

func NewSameVelocitySpawner(velocity float64, dim int) VelocitySpawner {
	return &SameVelocitySpawner{
		velocity: velocity,
		dim:      dim,
	}
}

func (s *SameVelocitySpawner) NextVelocity() vector.Vector {
	velocity := vector.ZeroVector(s.dim).SetDimension(1.0, 0)
	if s.dim == 2 {
		angle := math_util.RandFloat(-math.Pi, math.Pi)
		velocity = velocity.SetDimension(math.Sin(angle), 0).SetDimension(math.Cos(angle), 1)
	} else {
		angles := math_util.RandVectorByNumericBounaries(-math.Pi, math.Pi, s.dim-1)

		for i := range angles.Dimensions() {
			velocity = velocity.SetDimension(math.Cos(angles.Dimension(i)), i+1)
			for j := range i + 1 {
				velocity = velocity.SetDimension(velocity.Dimension(j)*math.Sin(angles.Dimension(i)), j)
			}
		}
	}

	return velocity.Mul(s.velocity)
}
