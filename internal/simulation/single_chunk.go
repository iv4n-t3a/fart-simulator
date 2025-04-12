package simulation

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type SingleChunkSimulation struct {
	time      float64
	chunk     chunk.Chunk
	reporters []metrics.Reporter
}

func NewSingleChunkSimulation(particlesAmount int) *SingleChunkSimulation {
	particles := particle.SpawnStaticRectangular2D(
		particlesAmount,
		vector.NewVector2D(0, 0),
		vector.NewVector2D(1, 1),
	)
	particle.AssignRandomSpeed2D(particles, 0, 5)

	rectContainerInst := container.NewRectContainer([]float64{1, 1})
	naiveChunk := chunk.NewNaiveChunk(1.0/100, rectContainerInst)

	naiveChunk.InitializeParticles(particles)

	return &SingleChunkSimulation{
		time:  0,
		chunk: naiveChunk,
	}
}

func (s *SingleChunkSimulation) SubscribeCollision(obs observers.CollisionObserver) {
	s.chunk.SubscribeCollision(obs)
}

func (s *SingleChunkSimulation) SubscribeCollisionWithContainer(obs observers.CollisionWithContainerObserver) {
	s.chunk.SubscribeCollisionWithContainer(obs)
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
