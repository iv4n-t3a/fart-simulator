package container

type ShrinkingRectContainer struct {
	SimpleRectContainer
	shrinkingSpeed float64
}

func NewShrinkingRectContainer(sides []float64, shrinkingSpeed float64) *ShrinkingRectContainer {
	return &ShrinkingRectContainer{
		SimpleRectContainer: SimpleRectContainer{sides: sides},
		shrinkingSpeed:      shrinkingSpeed,
	}
}

func (c *ShrinkingRectContainer) Tick(tick float64) {
	for i := range c.sides {
		c.sides[i] -= tick * c.shrinkingSpeed
	}
}
