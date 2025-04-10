package container

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type RectContainer struct {
	sides []float64
}

func (c *RectContainer) ProcessCollision(p *particle.Particle) bool {
  detectedCollision := false
  for i := range p.Pos.Dimensions() {
    if p.Pos.Dimension(i) <= 0 || p.Pos.Dimension(i) >= c.sides[i] {
      detectedCollision = true
      p.Vel.SetDimension(-p.Vel.Dimension(i), i)
    }
  }

  return detectedCollision
}
