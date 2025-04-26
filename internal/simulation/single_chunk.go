package simulation

import (
	"fmt"

	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/schollz/progressbar/v3"
)

type SingleChunkSimulation struct {
	time      float64
	chunk     chunk.Chunk
	particles []particle.Particle
}

func NewSingleChunkSimulation(particlesAmount int, container container.Container,
	chunkFactory chunk.ChunkFactory, spawner spawner.Spawner) *SingleChunkSimulation {
	particles := make([]particle.Particle, particlesAmount)

	for i := range particles {
		particles[i] = spawner.SpawnParticle()
	}
	chunk := chunkFactory.NewChunk(container, particles)

	return &SingleChunkSimulation{
		time:      0.0,
		chunk:     chunk,
		particles: particles,
	}
}

func (s *SingleChunkSimulation) Observers() *observers.ObserversComposition {
	return s.chunk.Observers()
}

func (s *SingleChunkSimulation) Run(time float64) {
	if time > 0 {
		s.runWithProgressBar(time)
	} else {
		s.runWithoutProgressBar(time)
	}
}

func (s *SingleChunkSimulation) runWithoutProgressBar(time float64) {
	for s.time < time || time < 0 {
		dt := s.chunk.EvaluateTimeStep()

		if dt <= 0 {
			panic(fmt.Sprintf("dt = %f is supposed to be positive", dt))
		}
		s.chunk.Simulate(dt)
		s.time += dt
	}
}

func (s *SingleChunkSimulation) runWithProgressBar(time float64) {
	bar := progressbar.Default(100)
	progressAdd := 0.0

	for s.time < time || time < 0 {
		dt := s.chunk.EvaluateTimeStep()

		if dt <= 0 {
			panic(fmt.Sprintf("dt = %f is supposed to be positive", dt))
		}
		s.chunk.Simulate(dt)
		s.time += dt
		progressAdd += dt

		for progressAdd >= time/100 {
			bar.Add(1)
			progressAdd -= time / 100
		}
	}
}

func (s *SingleChunkSimulation) Particles() []particle.Particle {
	return s.particles
}
