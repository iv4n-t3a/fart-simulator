package container

import (
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type HeatedRectContainer struct {
	RectContainer
	heatCoefficient float64
	timeSource      *metrics.TimeObserver
	threshold       float64
}

func NewHeatedRectContainer(sides []float64, heatCoefficient float64, timeSource *metrics.TimeObserver, threshold float64) *HeatedRectContainer {
	return &HeatedRectContainer{
		RectContainer:   RectContainer{sides: sides},
		heatCoefficient: heatCoefficient,
		timeSource:      timeSource,
		threshold:       threshold,
	}
}

func (c *HeatedRectContainer) ProcessCollision(p *particle.Particle) bool {
	detectedCollision := false
	for i := range p.Pos.Dimensions() {
		if p.Pos.Dimension(i) <= 0 || p.Pos.Dimension(i) >= c.sides[i] {
			detectedCollision = true
			p.Vel = p.Vel.SetDimension(-p.Vel.Dimension(i), i)
		}
	}

	if c.timeSource.Duration >= c.threshold {
		p.Vel = p.Vel.Mul(c.heatCoefficient)
	}

	return detectedCollision
}
