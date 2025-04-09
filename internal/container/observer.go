package container

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Observer interface {
	ParticleInserted(*particle.Particle)
	ParticleRemoved(*particle.Particle)
	Collision(*particle.Particle, *particle.Particle)
}
