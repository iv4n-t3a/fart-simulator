package metrics

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type CollisionWithContainerAggregatorObserver struct {
	timeSource *TimeObserver
	data       []ParticleData
}

func NewCollisionWithContainerAggregatorObserver(timeObserver *TimeObserver) *CollisionWithContainerAggregatorObserver {
	return &CollisionWithContainerAggregatorObserver{
		timeObserver,
		make([]ParticleData, 0),
	}
}

func (c *CollisionWithContainerAggregatorObserver) CollisionWithContainer(p *particle.Particle) {
	c.data = append(c.data, ParticleData{
		p.Pos.Coords(),
		p.Vel.Coords(),
		(*c.timeSource).Duration,
	})
}

func (c *CollisionWithContainerAggregatorObserver) Report() {
	fmt.Println("Reporting, len=", len(c.data))
	err := WriteParticleDataToBinary("data/collision_with_container_plot_bin", c.data)
	if err != nil {
		fmt.Println("Error writing particles data")
	}
}
