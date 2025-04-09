package box

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Box interface {
	AddParticle(particle.Particle)
	Subscribe(Observer)
	Simulate(delta_time float64)
	EvaluateTimeStep() float64
}
