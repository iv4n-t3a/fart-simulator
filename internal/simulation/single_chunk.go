package simulation

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
)

type SingleChunkSimulation struct {
	time      float64
	chunk     chunk.Chunk
	observer  *Observer
	reporters []*metrics.Reporter
}

func NewSingleChunkSimulation(chunk chunk.Chunk, observer *Observer) *SingleChunkSimulation {
	return &SingleChunkSimulation{
		time:     0,
		chunk:    chunk,
		observer: observer,
	}
}

func (s *SingleChunkSimulation) Subscribe(obs *Observer) {
	s.observer = obs
}

func (s *SingleChunkSimulation) AddReporter(reporter metrics.Reporter) {
	s.reporters = append(s.reporters, &reporter)
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
		(*s.reporters[i]).Report()
	}
}
