package velocity_spawner_test

import (
	"testing"

	"github.com/iv4n-t3a/fart-simulator/config"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"github.com/stretchr/testify/assert"
)

func TestSameVelocitySpawner2D(t *testing.T) {
	vel := 1.0
	s := velocity_spawner.NewSameVelocitySpawner(vel, 2)

	for range 100 {
		v := s.NextVelocity()
		assert.InEpsilon(t, vel, v.Length(), config.Eps, "Generated velocity differs from required")
	}
}

func TestSameVelocitySpawner3D(t *testing.T) {
	vel := 1.0
	s := velocity_spawner.NewSameVelocitySpawner(vel, 3)

	for range 100 {
		v := s.NextVelocity()
		assert.InEpsilon(t, vel, v.Length(), config.Eps, "Generated velocity differs from required")
	}
}
