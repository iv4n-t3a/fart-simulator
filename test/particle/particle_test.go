package particle_test

import (
	"testing"

	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"github.com/stretchr/testify/assert"
)

func TestParticlesCollisition(t *testing.T) {
	p1 := particle.Particle{
		Pos:    vector.NewVector3D(0.0, 0.0, 0.0),
		Vel:    vector.NewVector3D(1.0, 1.0, 0.0),
		Radius: 1.0,
		Mass:   1.0,
	}

	p2 := particle.Particle{
		Pos:    vector.NewVector3D(2.0, 0.0, 0.0),
		Vel:    vector.NewVector3D(-1.0, 0.0, 0.0),
		Radius: 1.0,
		Mass:   1.0,
	}

	particle.ProcessCollision(&p1, &p2)

	assert.Equal(t, -1.0, p1.Vel.X(), "Velocity check")
	assert.Equal(t, 1.0, p1.Vel.Y(), "Velocity check")
	assert.Equal(t, 0.0, p1.Vel.Z(), "Velocity check")

  assert.Equal(t, 1.0, p2.Vel.X(), "Velocity check")
	assert.Equal(t, 0.0, p2.Vel.Y(), "Velocity check")
	assert.Equal(t, 0.0, p2.Vel.Z(), "Velocity check")
}
