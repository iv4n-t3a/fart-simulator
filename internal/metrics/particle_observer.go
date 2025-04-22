package metrics

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type ParticleObserver struct {
	timeSource *TimeObserver
	data       []ParticleData
	dtWindow   float64
	lastUpdate float64
}

func NewParticleObserver(timeObserver *TimeObserver, dtWindow float64) *ParticleObserver {
	return &ParticleObserver{
		timeSource: timeObserver,
		data:       make([]ParticleData, 0),
		dtWindow:   dtWindow,
	}
}

func (p *ParticleObserver) ObserveParticle(particle *particle.Particle) {
	if p.timeSource.Duration-p.lastUpdate < p.dtWindow {
		return
	}
	p.lastUpdate = p.timeSource.Duration
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
