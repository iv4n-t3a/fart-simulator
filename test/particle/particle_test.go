package particle_test

import (
	"testing"

	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"github.com/stretchr/testify/assert"
)

func TestNoCollision(t *testing.T) {
	p1 := particle.Particle{
		Pos:    vector.NewVector3D(0.0, 0.0, 0.0),
		Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
		Radius: 1.0,
		Mass:   math_util.RandFloat(1.0, 10.0),
	}

	p2 := particle.Particle{
		Pos:    vector.NewVector3D(10.0, 10.0, 10.0),
		Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
		Radius: 1.0,
		Mass:   math_util.RandFloat(1.0, 10.0),
	}

	wasCollision := particle.ProcessCollision(&p1, &p2)
	assert.False(t, wasCollision, "Extra collision detected")
}

func TestParticlesCollisionStressSameMass(t *testing.T) {
	for range 100 {
		p1 := particle.Particle{
			Pos:    vector.NewVector3D(0.0, 0.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
			Radius: 1.0,
			Mass:   1.0,
		}

		p2 := particle.Particle{
			Pos:    vector.NewVector3D(1.0, 0.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
			Radius: 1.0,
			Mass:   1.0,
		}

		sumImpulsesBefore := p1.Impulse().Add(p2.Impulse())
		sumEnergyBefore := p1.KeeneticEnergy() + p2.KeeneticEnergy()
		vel1Before := p1.Vel
		vel2Before := p2.Vel

		wasCollision := particle.ProcessCollision(&p1, &p2)

		assert.True(t, wasCollision, "No collision detected")
		assert.InEpsilon(t, sumImpulsesBefore.Length(), p1.Impulse().Add(p2.Impulse()).Length(), config.Eps, "The law of conservation of impulse check")
		assert.InEpsilon(t, sumEnergyBefore, p1.KeeneticEnergy()+p2.KeeneticEnergy(), config.Eps, "The law of conservation of energy check")
		assert.NotEqual(t, vel1Before, p1.Vel, "Velocity not changed")
		assert.NotEqual(t, vel2Before, p2.Vel, "Velocity not changed")
	}
}

func TestParticlesCollisionStress(t *testing.T) {
	for range 100 {
		p1 := particle.Particle{
			Pos:    vector.NewVector3D(0.0, 0.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
			Radius: 1.0,
			Mass:   math_util.RandFloat(0.1, 1.0),
		}

		p2 := particle.Particle{
			Pos:    vector.NewVector3D(1.0, 0.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.1, 10.0, 3),
			Radius: 1.0,
			Mass:   math_util.RandFloat(0.1, 10.0),
		}

		sumImpulsesBefore := p1.Impulse().Add(p2.Impulse())
		sumEnergyBefore := p1.KeeneticEnergy() + p2.KeeneticEnergy()
		vel1Before := p1.Vel
		vel2Before := p2.Vel

		wasCollision := particle.ProcessCollision(&p1, &p2)

		assert.True(t, wasCollision, "No collision detected")
		assert.InEpsilon(t, sumImpulsesBefore.Length(), p1.Impulse().Add(p2.Impulse()).Length(), config.Eps, "The law of conservation of impulse check")
		assert.InEpsilon(t, sumEnergyBefore, p1.KeeneticEnergy()+p2.KeeneticEnergy(), config.Eps, "The law of conservation of energy check")
		assert.NotEqual(t, vel1Before, p1.Vel, "Velocity not changed")
		assert.NotEqual(t, vel2Before, p2.Vel, "Velocity not changed")
	}
}

func TestParticlesCollisionStress2D(t *testing.T) {
	for range 100 {
		p1 := particle.Particle{
			Pos:    vector.NewVector2D(0.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.0, 1000.0, 2),
			Radius: 1.0,
			Mass:   math_util.RandFloat(1.0, 10.0),
		}

		p2 := particle.Particle{
			Pos:    vector.NewVector2D(1.0, 0.0),
			Vel:    math_util.RandVectorByNumericBounaries(0.0, 1000.0, 2),
			Radius: 1.0,
			Mass:   math_util.RandFloat(1.0, 10.0),
		}

		sumImpulsesBefore := p1.Impulse().Add(p2.Impulse())
		sumEnergyBefore := p1.KeeneticEnergy() + p2.KeeneticEnergy()
		vel1Before := p1.Vel
		vel2Before := p2.Vel

		wasCollision := particle.ProcessCollision(&p1, &p2)

		assert.True(t, wasCollision, "No collision detected")
		assert.InEpsilon(t, sumImpulsesBefore.Length(), p1.Impulse().Add(p2.Impulse()).Length(), config.Eps, "The law of conservation of impulse check")
		assert.InEpsilon(t, sumEnergyBefore, p1.KeeneticEnergy()+p2.KeeneticEnergy(), config.Eps, "The law of conservation of energy check")
		assert.NotEqual(t, vel1Before, p1.Vel, "Velocity not changed")
		assert.NotEqual(t, vel2Before, p2.Vel, "Velocity not changed")
	}
}
