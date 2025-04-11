package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type Chunk interface {
	AddParticle(particle.Particle)

	SubscribeParticleInserted(*observers.ParticleInsertedObserver)
	SubscribeParticleRemoved(*observers.ParticleRemovedObserver)
	SubscribeCollision(*observers.CollisionObserver)
	SubscribeCollisionWithContainer(*observers.CollisionWithContainerObserver)
	EvaluateTimeStep() float64

	// Mustn't be called on values greater than EvaluateTimeStep
	Simulate(deltaTime float64)
}
