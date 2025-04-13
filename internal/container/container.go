package container

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Container interface {
	ProcessCollision(*particle.Particle) bool
  TimeBeforeCollision(particle.Particle) float64
}
