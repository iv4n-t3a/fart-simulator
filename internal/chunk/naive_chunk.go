package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type NaiveChunk struct {
	dt        float64
	particles []particle.Particle
	container container.Container
}

func NewNaiveChunk(dt float64) *NaiveChunk {
	return &NaiveChunk{dt: dt}
}

func (c *NaiveChunk) AddParticle(p particle.Particle) {
	c.particles = append(c.particles, p)
}

func (c *NaiveChunk) Subscribe(*Observer) {
	panic("Unimplemented")
}

func (c *NaiveChunk) Simulate(dt float64) {
	// Update positions
	for i := range c.particles {
		p := &c.particles[i]
		p.Pos.Add(p.Vel.Mul(dt))
	}

	for i := range c.particles {
		c.container.ProcessCollision(&c.particles[i])
		for j := range c.particles[i:] {
			particle.ProcessCollision(&c.particles[i], &c.particles[j])
		}
	}
}

func (c *NaiveChunk) EvaluateTimeStep() float64 {
	return c.dt
}
