package spawner

import (
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type RectSpawner struct {
	maxVelocity float64
	sides       []float64
}

func NewRectSpawner(maxVelocity float64, c container.HeatedRectContainer) Spawner {
	return &RectSpawner{
		maxVelocity: maxVelocity,
		sides:       c.GetSides(),
	}
}

func (s *RectSpawner) SpawnParticle() particle.Particle {
	return particle.Particle{
		Pos:    math_util.RandVectorByVectorBounaries(vector.ZeroVector(len(s.sides)), vector.NewVector(s.sides)),
		Vel:    math_util.RandVectorByNumericBounaries(-s.maxVelocity, s.maxVelocity, len(s.sides)),
		Radius: config.Radius,
		Mass:   config.Mass,
	}
}
