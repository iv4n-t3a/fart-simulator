package particle

import (
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

func SpawnStaticRectangular2D(amount int, min vector.Vector2D, max vector.Vector2D) *[]Particle {
	particles := make([]Particle, amount)
	for i := range particles {
		particles[i] = Particle{
			vector.NewVector2D(math_util.RandFloat(min.X(), max.X()), math_util.RandFloat(min.Y(), max.Y())),
			vector.NewVector2D(0, 0),
			config.Radius,
			config.Mass,
		}
	}
	return &particles
}

func AssignRandomSpeed2D(particles *[]Particle, min float64, max float64) {
	for i := range *particles {
		(*particles)[i].Vel = vector.NewVector2D(math_util.RandFloat(min, max), math_util.RandFloat(min, max))
	}
}
