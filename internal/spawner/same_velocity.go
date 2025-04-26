package spawner

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type SameVelocitySpawner struct {
	velocity  float64
	sides     []float64
	nextIndex int64
}

func NewSameVelocitySpawner(velocity float64, c container.RectContainer) Spawner {
	return &SameVelocitySpawner{
		velocity:  velocity,
		sides:     c.GetSides(),
		nextIndex: 0,
	}
}

func (s *SameVelocitySpawner) SpawnParticle() particle.Particle {
	velocity := vector.ZeroVector(len(s.sides)).SetDimension(1.0, 0)
	if len(s.sides) == 2 {
		angle := math_util.RandFloat(-math.Pi, math.Pi)
		velocity = velocity.SetDimension(math.Sin(angle), 0).SetDimension(math.Cos(angle), 1)
	} else {
		angles := math_util.RandVectorByNumericBounaries(-math.Pi, math.Pi, len(s.sides)-1)

		for i := range angles.Dimensions() {
			velocity = velocity.SetDimension(math.Cos(angles.Dimension(i)), i+1)
			for j := range i + 1 {
				velocity = velocity.SetDimension(velocity.Dimension(j)*math.Sin(angles.Dimension(i)), j)
			}
		}
	}
	s.nextIndex += 1

	return particle.Particle{
		Pos:    math_util.RandVectorByVectorBounaries(vector.ZeroVector(len(s.sides)), vector.NewVector(s.sides)),
		Vel:    velocity.Mul(s.velocity),
		Radius: config.Radius,
		Mass:   config.Mass,
		Index:  s.nextIndex - 1,
	}
}
