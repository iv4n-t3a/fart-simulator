package container

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type SimpleRectContainer struct {
	sides []float64
}

func NewSimpleRectContainer(sides []float64) *SimpleRectContainer {
	return &SimpleRectContainer{sides: sides}
}

func (c *SimpleRectContainer) ProcessCollision(p *particle.Particle) bool {
	detectedCollision := false
	for i := range p.Pos.Dimensions() {
		if p.Pos.Dimension(i) <= 0 || p.Pos.Dimension(i) >= c.sides[i] {
			detectedCollision = true
			if p.Pos.Dimension(i) <= 0 {
				p.Vel = p.Vel.SetDimension(math.Abs(p.Vel.Dimension(i)), i)
			} else {
				p.Vel = p.Vel.SetDimension(-math.Abs(p.Vel.Dimension(i)), i)
			}
		}
	}

	return detectedCollision
}

func (c *SimpleRectContainer) TimeBeforeCollision(p particle.Particle) float64 {
	res := math.Inf(1)

	for i := range c.sides {
		v := p.Vel.Dimension(i)
		x := p.Pos.Dimension(i)

		if x < p.Radius || x >= c.sides[i]-p.Radius {
			return 0.0
		}

		if v > 0 {
			res = min((c.sides[i]-x-p.Radius)/v, res)
		} else {
			res = min((p.Radius-x)/v, res)
		}
	}

	return res
}

func (c *SimpleRectContainer) GetSides() []float64 {
	return c.sides
}
