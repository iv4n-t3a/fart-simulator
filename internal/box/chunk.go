package box

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Chunk interface {
	AddParticle(particle.Particle)
	Subscribe(*Observer)
	Simulate(deltaTime float64)
	EvaluateTimeStep() float64
}
