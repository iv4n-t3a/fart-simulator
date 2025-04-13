package container

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type RectContainer struct {
	sides []float64
}

func NewRectContainer(sides []float64) *RectContainer {
	return &RectContainer{sides: sides}
}

func (c *RectContainer) ProcessCollision(p *particle.Particle) bool {
	detectedCollision := false
	for i := range p.Pos.Dimensions() {
		if p.Pos.Dimension(i) <= 0 || p.Pos.Dimension(i) >= c.sides[i] {
			detectedCollision = true
			p.Vel = p.Vel.SetDimension(-p.Vel.Dimension(i), i)
		}
	}

	return detectedCollision
}

func (c *RectContainer) TimeBeforeCollision(p particle.Particle) float64 {
	res := math.Inf(1)

	for i := range c.sides {
		v := p.Vel.Dimension(i)
		x := p.Pos.Dimension(i)

		if x < p.Radius || x >= c.sides[i] - p.Radius {
			return 0.0
		}

		if v > 0 {
			res = min((c.sides[i] - x - p.Radius) / v)
		} else {
			res = min((p.Radius-x)/v, res)
		}
	}

	return res
}

func (c RectContainer) GetSides() []float64 {
	return c.sides
}
