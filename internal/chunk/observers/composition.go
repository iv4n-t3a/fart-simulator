package observers

import (
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type ObserversComposition struct {
	particleInsertedObservers       []ParticleInsertedObserver
	particleRemovedObservers        []ParticleRemovedObserver
	collisionObservers              []CollisionObserver
	collisionWithContainerObservers []CollisionWithContainerObserver
	timeObservers                   []TimeObserver
	particleObservers               []ParticleObserver
}

func (c *ObserversComposition) SubscribeParticleInserted(obs ParticleInsertedObserver) {
	c.particleInsertedObservers = append(c.particleInsertedObservers, obs)
}

func (c *ObserversComposition) SubscribeParticleRemoved(obs ParticleRemovedObserver) {
	c.particleRemovedObservers = append(c.particleRemovedObservers, obs)
}

func (c *ObserversComposition) SubscribeCollision(obs CollisionObserver) {
	c.collisionObservers = append(c.collisionObservers, obs)
}

func (c *ObserversComposition) SubscribeCollisionWithContainer(obs CollisionWithContainerObserver) {
	c.collisionWithContainerObservers = append(c.collisionWithContainerObservers, obs)
}

func (c *ObserversComposition) SubscribeTime(obs TimeObserver) {
	c.timeObservers = append(c.timeObservers, obs)
}

func (c *ObserversComposition) SubscribeParticle(obs ParticleObserver) {
	c.particleObservers = append(c.particleObservers, obs)
}

func (c *ObserversComposition) ParticleInserted(p *particle.Particle) {
	for i := range c.particleInsertedObservers {
		c.particleInsertedObservers[i].ParticleInserted(p)
	}
}

func (c *ObserversComposition) ParticleRemoved(p *particle.Particle) {
	for i := range c.particleRemovedObservers {
		c.particleRemovedObservers[i].ParticleRemoved(p)
	}
}

func (c *ObserversComposition) Collision(p1 *particle.Particle, p2 *particle.Particle) {
	for i := range c.collisionObservers {
		c.collisionObservers[i].Collision(p1, p2)
	}
}

func (c *ObserversComposition) CollisionWithContainer(p *particle.Particle) {
	for i := range c.collisionWithContainerObservers {
		c.collisionWithContainerObservers[i].CollisionWithContainer(p)
	}
}

func (c *ObserversComposition) Tick(t float64) {
	for i := range c.timeObservers {
		c.timeObservers[i].Tick(t)
	}
}

func (c *ObserversComposition) ObserveParticle(p *particle.Particle) {
	for i := range c.particleObservers {
		c.particleObservers[i].ObserveParticle(p)
	}
}
