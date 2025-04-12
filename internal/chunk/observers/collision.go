package observers

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type CollisionObserver interface {
	Collision(*particle.Particle, *particle.Particle)
}
