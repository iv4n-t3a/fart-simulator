package spawner_test

import (
	"testing"

	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/stretchr/testify/assert"
)

func TestSameVelocitySpawner2D(t *testing.T) {
	vel := 1.0
	cont := container.NewRectContainer([]float64{1.0, 1.0})
	s := spawner.NewSameVelocitySpawner(vel, *cont)

	for range 100 {
		p := s.SpawnParticle()
		assert.InEpsilon(t, vel, p.Vel.Length(), config.Eps, "Generated velocity differs from required")
	}
}

func TestSameVelocitySpawner3D(t *testing.T) {
	vel := 1.0
	cont := container.NewRectContainer([]float64{1.0, 1.0, 1.0})
	s := spawner.NewSameVelocitySpawner(vel, *cont)

	for range 100 {
		p := s.SpawnParticle()
		assert.InEpsilon(t, vel, p.Vel.Length(), config.Eps, "Generated velocity differs from required")
	}
}
