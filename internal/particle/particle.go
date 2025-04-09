package particle

import "github.com/iv4n-t3a/fart-simulator/internal/vector"

type Particle struct {
	pos    vector.Vector
	vel    vector.Vector
	radius float64
}
