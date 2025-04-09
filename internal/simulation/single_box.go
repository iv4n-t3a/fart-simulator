package simulation

import "github.com/iv4n-t3a/fart-simulator/internal/box"

type SingleBoxSimulation struct {
	time      float64
	container box.Chunk
	observer  *Observer
}

func (s *SingleBoxSimulation) Subscribe(obs *Observer) {
	s.observer = obs
}

func (s *SingleBoxSimulation) Run(time float64) {
	for s.time < time {
		dt := s.container.EvaluateTimeStep()
		s.container.Simulate(dt)
		time += dt
	}
}
