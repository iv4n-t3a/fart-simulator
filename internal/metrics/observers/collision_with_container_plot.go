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

type CollisionWithContainerPlotObserver struct {
	timeSource *TimeObserver
	data       []CollisionData
}

func NewCollisionWithContainerPlotObserver(timeObserver *TimeObserver) *CollisionWithContainerPlotObserver {
	return &CollisionWithContainerPlotObserver{
		timeObserver,
		make([]CollisionData, 0),
	}
}

func (c *CollisionWithContainerPlotObserver) CollisionWithContainer(p *particle.Particle) {
	c.data = append(c.data, CollisionData{
		p.Pos.Coords(),
		p.Vel.Coords(),
		(*c.timeSource).Duration,
	})
}

func (c *CollisionWithContainerPlotObserver) Report() {
	jsonData, err := json.MarshalIndent(c.data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling collision data")
	}
	err = os.WriteFile("internal/metrics/files/collision_with_container_plot.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing collision data")
	}
}
