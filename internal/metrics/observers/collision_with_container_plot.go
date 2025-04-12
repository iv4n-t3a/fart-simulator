package observers

import (
	"encoding/json"
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"os"
)

type CollisionData struct {
	Position []float64 `json:"position"`
	Velocity []float64 `json:"velocity"`
	Time     float64   `json:"time"`
}

type CollisionWithContainerAggregatorObserver struct {
	timeSource *TimeObserver
	data       []CollisionData
}

func NewCollisionWithContainerAggregatorObserver(timeObserver *TimeObserver) *CollisionWithContainerAggregatorObserver {
	return &CollisionWithContainerAggregatorObserver{
		timeObserver,
		make([]CollisionData, 0),
	}
}

func (c *CollisionWithContainerAggregatorObserver) CollisionWithContainer(p *particle.Particle) {
	c.data = append(c.data, CollisionData{
		p.Pos.Coords(),
		p.Vel.Coords(),
		(*c.timeSource).Duration,
	})
}

func (c *CollisionWithContainerAggregatorObserver) Report() {
	jsonData, err := json.MarshalIndent(c.data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling collision data")
	}
	err = os.WriteFile("data/collision_with_container_plot.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing collision data")
	}
}
