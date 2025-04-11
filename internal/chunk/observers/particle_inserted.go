package observers

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type ParticleInsertedObserver interface {
	ParticleInserted(particle *particle.Particle)
}
