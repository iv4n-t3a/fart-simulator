package metrics

import (
	"encoding/json"
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"os"
)

type ParticleObserver struct {
	timeSource *TimeObserver
	data       []CollisionData
}

func NewParticleObserver(timeObserver *TimeObserver) *ParticleObserver {
	return &ParticleObserver{
		timeSource: timeObserver,
		data:       make([]CollisionData, 0),
	}
}

func (p *ParticleObserver) ObserveParticle(particle *particle.Particle) {
	p.data = append(p.data, CollisionData{
		particle.Pos.Coords(),
		particle.Vel.Coords(),
		(*p.timeSource).Duration,
	})
}

func (p *ParticleObserver) Report() {
	jsonData, err := json.MarshalIndent(p.data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling particles data")
	}
	err = os.WriteFile("data/particles_data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing particles data")
	}
}
