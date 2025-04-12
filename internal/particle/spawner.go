package particle

import (
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"math/rand"
)

func SpawnStaticRectangular2D(amount int, min vector.Vector2D, max vector.Vector2D) *[]Particle {
	particles := make([]Particle, amount)
	for i := range particles {
		particles[i] = Particle{
			vector.NewVector2D(min.X()+rand.Float64()*(max.X()-min.X()), min.Y()+rand.Float64()*(max.Y()-min.Y())),
			vector.NewVector2D(0, 0),
			config.Radius,
			config.Mass,
		}
	}
	return &particles
}
