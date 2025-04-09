package simulation

import "github.com/iv4n-t3a/fart-simulator/internal/chunk"

type SingleChunkSimulation struct {
	time     float64
	chunk    chunk.Chunk
	observer *Observer
}

func (s *SingleChunkSimulation) Subscribe(obs *Observer) {
	s.observer = obs
}

func (s *SingleChunkSimulation) Run(time float64) {
	for s.time < time {
		dt := s.chunk.EvaluateTimeStep()
		s.chunk.Simulate(dt)
		time += dt
	}
}
