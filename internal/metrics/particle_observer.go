package metrics

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type ParticleObserver struct {
	timeSource *TimeObserver
	data       []ParticleData
}

func NewParticleObserver(timeObserver *TimeObserver) *ParticleObserver {
	return &ParticleObserver{
		timeSource: timeObserver,
		data:       make([]ParticleData, 0),
	}
}

func (p *ParticleObserver) ObserveParticle(particle *particle.Particle) {
	p.data = append(p.data, ParticleData{
		particle.Pos.Coords(),
		particle.Vel.Coords(),
		(*p.timeSource).Duration,
	})
}

func (p *ParticleObserver) Report() {
	err := WriteParticleDataToBinary("data/particles_data_bin", p.data)
	if err != nil {
		fmt.Println("Error writing particles data")
	}
}
