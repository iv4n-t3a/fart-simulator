package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type NaiveChunk struct {
	dt                              float64
	particles                       []particle.Particle
	container                       container.Container
	particleInsertedObservers       []*observers.ParticleInsertedObserver
	particleRemovedObservers        []*observers.ParticleRemovedObserver
	collisionObservers              []*observers.CollisionObserver
	collisionWithContainerObservers []*observers.CollisionWithContainerObserver
}

func NewNaiveChunk(dt float64, container container.Container) *NaiveChunk {
	return &NaiveChunk{
		dt:        dt,
		container: container,
	}
}

func (c *NaiveChunk) AddParticle(p particle.Particle) {
	for i := range c.particleInsertedObservers {
		(*c.particleInsertedObservers[i]).ParticleInserted(&p)
	}
	c.particles = append(c.particles, p)
}

func (c *NaiveChunk) InitializeParticles(particles *[]particle.Particle) {
	c.particles = *particles
}

func (c *NaiveChunk) SubscribeParticleInserted(obs observers.ParticleInsertedObserver) {
	c.particleInsertedObservers = append(c.particleInsertedObservers, &obs)
}

func (c *NaiveChunk) SubscribeParticleRemoved(obs observers.ParticleRemovedObserver) {
	c.particleRemovedObservers = append(c.particleRemovedObservers, &obs)
}

func (c *NaiveChunk) SubscribeCollision(obs observers.CollisionObserver) {
	c.collisionObservers = append(c.collisionObservers, &obs)
}

func (c *NaiveChunk) SubscribeCollisionWithContainer(obs observers.CollisionWithContainerObserver) {
	c.collisionWithContainerObservers = append(c.collisionWithContainerObservers, &obs)
}

func (c *NaiveChunk) Simulate(dt float64) {
	// Update positions
	for i := range c.particles {
		p := &c.particles[i]
		p.Pos = p.Pos.Add(p.Vel.Mul(dt))
	}

	for i := range c.particles {
		if c.container.ProcessCollision(&c.particles[i]) {
			for i := range c.collisionWithContainerObservers {
				(*c.collisionWithContainerObservers[i]).CollisionWithContainer(&c.particles[i])
			}
		}
		for j := i + 1; j < len(c.particles); j++ {
			if particle.ProcessCollision(&c.particles[i], &c.particles[j]) {
				for k := range c.collisionObservers {
					(*c.collisionObservers[k]).Collision(&c.particles[i], &c.particles[j])
				}
			}
		}
	}
}

func (c *NaiveChunk) EvaluateTimeStep() float64 {
	return c.dt
}
