package container

import (
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type Container interface {
  IsInside(vector.Vector) bool
	ProcessCollision(*particle.Particle) bool
	TimeBeforeCollision(particle.Particle) float64
}
