package spawner

import (
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
)

type SpawnerImpl struct {
	radius    float64
	mass      float64
	nextIndex int64
	pos       position_spawner.PositionSpawner
	vel       velocity_spawner.VelocitySpawner
}

func NewSpawnerImpl(radius float64, mass float64, container container.Container,
	pos position_spawner.PositionGenerator, vel velocity_spawner.VelocitySpawner) Spawner {
	return &SpawnerImpl{
		radius:    radius,
		mass:      mass,
		nextIndex: 0,
		pos:       *position_spawner.NewPositionSpawner(pos, container),
		vel:       vel,
	}
}

func (s *SpawnerImpl) SpawnParticle() particle.Particle {
	s.nextIndex += 1
	return particle.Particle{
		Pos:    s.pos.NextPosition(s.radius),
		Vel:    s.vel.NextVelocity(),
		Radius: s.radius,
		Mass:   s.mass,
		Index:  s.nextIndex - 1,
	}
}
