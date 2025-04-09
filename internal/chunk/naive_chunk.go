package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type NaiveChunk struct {
	dt        float64
	particles []particle.Particle
	container container.Container
	observers []Observer
}

func NewNaiveChunk(dt float64) *NaiveChunk {
	return &NaiveChunk{dt: dt}
}

func (c *NaiveChunk) AddParticle(p particle.Particle) {
	for i := range c.observers {
		c.observers[i].ParticleInserted(&p)
	}
	c.particles = append(c.particles, p)
}

func (c *NaiveChunk) Subscribe(obs Observer) {
	c.observers = append(c.observers, obs)
}

func (c *NaiveChunk) Simulate(dt float64) {
	// Update positions
	for i := range c.particles {
		p := &c.particles[i]
		p.Pos.Add(p.Vel.Mul(dt))
	}

	for i := range c.particles {
		if c.container.ProcessCollision(&c.particles[i]) {
			for i := range c.observers {
				c.observers[i].CollisionWithContainer(&c.particles[i])
			}
		}
    for j := i+1; j < len(c.particles); j++ {
			if particle.ProcessCollision(&c.particles[i], &c.particles[j]) {
				for k := range c.observers {
					c.observers[k].Collision(&c.particles[i], &c.particles[j])
				}
			}
		}
	}
}

func (c *NaiveChunk) EvaluateTimeStep() float64 {
	return c.dt
}
