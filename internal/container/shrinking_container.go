package container

import (
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"math"
)

type ShrinkingRectContainer struct {
	SimpleRectContainer
	shrinkingSpeed float64
	resistance     float64
}

func NewShrinkingRectContainer(sides []float64, shrinkingSpeed float64, resistance float64) *ShrinkingRectContainer {
	return &ShrinkingRectContainer{
		SimpleRectContainer: SimpleRectContainer{sides: sides},
		shrinkingSpeed:      shrinkingSpeed,
		resistance:          resistance,
	}
}

func (c *ShrinkingRectContainer) Tick(tick float64) {
	for i := range c.sides {
		c.sides[i] -= tick * c.shrinkingSpeed
	}
}

func (c *ShrinkingRectContainer) ProcessCollision(p *particle.Particle) bool {
	detectedCollision := false
	for i := range p.Pos.Dimensions() {
		if p.Pos.Dimension(i) <= 0 || p.Pos.Dimension(i) >= c.sides[i] {
			detectedCollision = true
			if p.Pos.Dimension(i) <= 0 {
				p.Vel = p.Vel.SetDimension(math.Abs(p.Vel.Dimension(i)), i)
			} else {
				p.Vel = p.Vel.SetDimension(-math.Abs(p.Vel.Dimension(i)), i)
			}
			p.Vel = p.Vel.SetDimension(p.Vel.Dimension(i)+(c.sides[i]-p.Pos.Dimension(i))*c.resistance, i)
		}
	}

	return detectedCollision
}
