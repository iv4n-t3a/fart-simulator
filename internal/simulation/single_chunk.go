package simulation

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/schollz/progressbar/v3"
)

type SingleChunkSimulation struct {
	time  float64
	chunk chunk.Chunk
}

func NewSingleChunkSimulation(particlesAmount int, container container.Container,
	chunkFactory chunk.ChunkFactory, spawner spawner.Spawner) *SingleChunkSimulation {
	particles := make([]particle.Particle, particlesAmount)

	for i := range particles {
		particles[i] = spawner.SpawnParticle()
	}
	chunk := chunkFactory.NewChunk(container, particles)

	return &SingleChunkSimulation{
		time:  0.0,
		chunk: chunk,
	}
}

func (s *SingleChunkSimulation) Observers() *observers.ObserversComposition {
	return s.chunk.Observers()
}

func (s *SingleChunkSimulation) Run(time float64) {
	bar := progressbar.Default(100)
	progressAdd := 0.0

	for s.time < time {
		dt := s.chunk.EvaluateTimeStep()
		s.chunk.Simulate(dt)
		s.time += dt
		progressAdd += dt

		if progressAdd >= time/100 {
			bar.Add(1)
			progressAdd -= time / 100
		}
	}
}
