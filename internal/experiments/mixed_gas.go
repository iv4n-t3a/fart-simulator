package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunMixedGasSimulation(dim int) {
	side := 0.05
	dt := 1e-7
	startVelocity := 1.0
	radius1 := 0.0001
	radius2 := 0.0002
	mass1 := 1.0
	mass2 := 1.5
	count := 100

	sides := make([]float64, dim)

	for i := range sides {
		sides[i] = side
	}

	velSpawner := velocity_spawner.NewMoveOneVelocitySpawner(startVelocity, len(sides))

	containerInst := container.NewSimpleRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(dt)

	spawnerInst := spawner.NewSeparatedPositionSpawner(
		velSpawner,
		velSpawner,
		0.5,
		0.5,
		containerInst,
		mass1,
		mass2,
		radius1,
		radius2,
	)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)
	visualisationInst := visualisation.StartVisualisation(dim)

	visualisationInst.Init(simulationInst.Particles())
	simulationInst.Observers().SubscribeParticle(visualisationInst)
	simulationInst.Observers().SubscribeCollision(visualisationInst)
	simulationInst.Observers().SubscribeCollisionWithContainer(visualisationInst)

	defer visualisationInst.Report()
	simulationInst.Run(-1.0)
}
