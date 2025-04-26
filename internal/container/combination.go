package container

import (
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type ContainerCombination struct {
	container1 Container
	container2 Container
}

func Pair(c2 Container, c1 Container) *ContainerCombination {
	return &ContainerCombination{
		container1: c1,
		container2: c2,
	}
}

func (c *ContainerCombination) IsInside(v vector.Vector) bool {
	return c.container1.IsInside(v) && c.container2.IsInside(v)
}

func (c *ContainerCombination) ProcessCollision(p *particle.Particle) bool {
	return c.container1.ProcessCollision(p) || c.container2.ProcessCollision(p)
}

func (c *ContainerCombination) TimeBeforeCollision(p particle.Particle) float64 {
	return min(c.container1.TimeBeforeCollision(p), c.container2.TimeBeforeCollision(p))
}
