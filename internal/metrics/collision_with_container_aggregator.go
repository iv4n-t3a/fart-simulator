package metrics

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type CollisionWithContainerCounterObserver struct {
	counter uint64
}

func NewCollisionWithContainerCounterObserver() *CollisionWithContainerCounterObserver {
	return &CollisionWithContainerCounterObserver{}
}

func (c *CollisionWithContainerCounterObserver) CollisionWithContainer(*particle.Particle) {
	c.counter++
}

func (c *CollisionWithContainerCounterObserver) Report() {
	fmt.Printf("collision with container counter %d\n", c.counter)
}
