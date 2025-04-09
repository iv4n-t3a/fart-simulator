package chunk

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Chunk interface {
	AddParticle(particle.Particle)
	Subscribe(*Observer)
	EvaluateTimeStep() float64

	// Mustn't be called on values greater than EvaluateTimeStep
	Simulate(deltaTime float64)
}
