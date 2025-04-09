package simulation

import "github.com/iv4n-t3a/fart-simulator/internal/box"

type SingleBoxSimulation struct {
	time      float64
	container box.Box
	observers []Observer
}

func (s *SingleBoxSimulation) Subscribe(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *SingleBoxSimulation) Run(time float64) {
	for s.time < time {
		dt := s.container.EvaluateTimeStep()
		s.container.Simulate(dt)
		time += dt
	}
}
