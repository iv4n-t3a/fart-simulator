package observers

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type ParticleRemovedObserver interface {
	ParticleRemoved(particle *particle.Particle)
}
