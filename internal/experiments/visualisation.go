package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunVisualisation(dim int) {
	side := 0.05
	dt := 1e-7
	max_velocity := 1.0
	radius := 0.0001
	mass := 1.0
	count := 100

	sides := make([]float64, dim)

	for i := range sides {
		sides[i] = side
	}

	containerInst := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(dt)
	spawnerInst := spawner.NewRectSpawner(max_velocity, radius, mass, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

	visualisation := visualisation.StartVisualisation(dim)

	simulationInst.Observers().SubscribeParticle(visualisation)
	simulationInst.Observers().SubscribeCollision(visualisation)
	simulationInst.Observers().SubscribeCollisionWithContainer(visualisation)

	defer visualisation.Report()

	simulationInst.Run(-1.0)
}
