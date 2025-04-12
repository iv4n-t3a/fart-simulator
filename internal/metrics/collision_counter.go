package metrics

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type CollisionCounterObserver struct {
	counter uint64
}

func NewCollisionCounterObserver() *CollisionCounterObserver {
	return &CollisionCounterObserver{}
}

func (c *CollisionCounterObserver) Collision(*particle.Particle, *particle.Particle) {
	c.counter++
}

func (c *CollisionCounterObserver) Report() {
	fmt.Printf("collision counter %d\n", c.counter)
}
