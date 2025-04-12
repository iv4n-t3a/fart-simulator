package simulation

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

type SingleChunkSimulation struct {
	time      float64
	chunk     chunk.Chunk
	reporters []metrics.Reporter
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

func (s *SingleChunkSimulation) SubscribeCollision(obs observers.CollisionObserver) {
	s.chunk.SubscribeCollision(obs)
}

func (s *SingleChunkSimulation) SubscribeCollisionWithContainer(obs observers.CollisionWithContainerObserver) {
	s.chunk.SubscribeCollisionWithContainer(obs)
}

func (s *SingleChunkSimulation) SubscribeTime(obs observers.TimeObserver) {
	s.chunk.SubscribeTime(obs)
}

func (s *SingleChunkSimulation) AddReporter(reporter metrics.Reporter) {
	s.reporters = append(s.reporters, reporter)
}

func (s *SingleChunkSimulation) Run(time float64) {
	for s.time < time {
		dt := s.chunk.EvaluateTimeStep()
		s.chunk.Simulate(dt)
		s.time += dt
	}
}

func (s *SingleChunkSimulation) ReportMetrics() {
	for i := range s.reporters {
		s.reporters[i].Report()
	}
}
