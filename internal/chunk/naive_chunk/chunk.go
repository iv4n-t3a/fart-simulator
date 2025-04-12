package naive_chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type NaiveChunk struct {
	dt        float64
	particles []particle.Particle
	container container.Container
	observers observers.ObserversComposition
}

func newNaiveChunk(dt float64, container container.Container, particles []particle.Particle) *NaiveChunk {
	return &NaiveChunk{
		dt:        dt,
		container: container,
		particles: particles,
	}
}

func (c *NaiveChunk) AddParticle(p particle.Particle) {
  c.observers.ParticleInserted(&p)
	c.particles = append(c.particles, p)
}

func (c *NaiveChunk) Simulate(dt float64) {
	c.observers.Tick(dt)

	// Update positions
	for i := range c.particles {
		p := &c.particles[i]
		p.Pos = p.Pos.Add(p.Vel.Mul(dt))
	}

	for i := range c.particles {
		if c.container.ProcessCollision(&c.particles[i]) {
			c.observers.CollisionWithContainer(&c.particles[i])
		}
		for j := i + 1; j < len(c.particles); j++ {
			if particle.ProcessCollision(&c.particles[i], &c.particles[j]) {
				c.observers.Collision(&c.particles[i], &c.particles[j])
			}
		}
	}
}

func (c *NaiveChunk) EvaluateTimeStep() float64 {
	return c.dt
}

func (c *NaiveChunk) Observers() *observers.ObserversComposition {
	return &c.observers
}
