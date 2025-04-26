package spawner

import (
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type MoveSomeSpawner struct {
	innerSpawner Spawner
	count        int
	sides        []float64
	nextIndex    int64
}

func NewMoveSomeSpawner(velocity float64, count int, c *container.RectContainer) Spawner {
	return &MoveSomeSpawner{
		innerSpawner: NewSameVelocitySpawner(velocity, *c),
		count:        count,
		sides:        c.GetSides(),
		nextIndex:    0,
	}
}

func (s *MoveSomeSpawner) SpawnParticle() particle.Particle {
	s.nextIndex += 1
	if s.count > 0 {
		s.count -= 1
		return s.innerSpawner.SpawnParticle()
	}
	return particle.Particle{
		Pos:    math_util.RandVectorByVectorBounaries(vector.ZeroVector(len(s.sides)), vector.NewVector(s.sides)),
		Vel:    vector.ZeroVector(len(s.sides)),
		Radius: config.Radius,
		Mass:   config.Mass,
		Index:  s.nextIndex - 1,
	}
}
