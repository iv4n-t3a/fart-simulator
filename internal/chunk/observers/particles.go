package observers

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type ParticleObserver interface {
	ObserveParticle(particle *particle.Particle)
}
