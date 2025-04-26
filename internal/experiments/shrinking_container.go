package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunShrinkingContainerSimulation(dim int) {
	initialSide := 0.1
	dt := 1e-7
	count := 100
	mass := 1.0
	radius := 0.0001
	maxVelocity := 1.0

	sides := make([]float64, dim)

	for i := range sides {
		sides[i] = initialSide
	}

	containerInst := container.NewShrinkingRectContainer(sides, 0.1)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(dt)
	spawnerInst := spawner.NewRectSpawner(maxVelocity, radius, mass, containerInst)
	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

	visualisationInst := visualisation.StartVisualisation(dim)

	simulationInst.Observers().SubscribeParticle(visualisationInst)
	simulationInst.Observers().SubscribeCollision(visualisationInst)
	simulationInst.Observers().SubscribeCollisionWithContainer(visualisationInst)

	simulationInst.Observers().SubscribeTime(containerInst)

	defer visualisationInst.Report()

	simulationInst.Run(-1.0)
}
