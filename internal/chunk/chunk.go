package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type Chunk interface {
	AddParticle(particle.Particle)
	EvaluateTimeStep() float64
	Simulate(deltaTime float64) // Mustn't be called on values greater than EvaluateTimeStep
	Observers() *observers.ObserversComposition
}
