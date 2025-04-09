package simulation

import "github.com/iv4n-t3a/fart-simulator/internal/container"

type SingleContainerSimulation struct {
	time      float64
	container container.Container
	observers []Observer
}

func (s *SingleContainerSimulation) Subscribe(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *SingleContainerSimulation) Run(time float64) {
	for s.time < time {
		dt := s.container.EvaluateTimeStep()
		s.container.Simulate(dt)
		time += dt
	}
}
