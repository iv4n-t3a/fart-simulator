package particle

import "github.com/iv4n-t3a/fart-simulator/internal/vector"

type Particle struct {
	Pos    vector.Vector
	Vel    vector.Vector
	Radius float64
}

func ProcessCollision(a *Particle, b *Particle) {
  if a.Pos.Dist(b.Pos) < a.Radius + b.Radius {
    return
  }

}
