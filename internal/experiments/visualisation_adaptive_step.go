package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunVisualisationWithAdaptiveStep(dim int) {
	side := 0.05
	// dt := 1e-7
	startVelocity := 1.0
	count := 1000
  radius := 1e-4
  mass := 1.5e-20

	sides := make([]float64, dim)

	for i := range sides {
		sides[i] = side
	}

	velSpawner := velocity_spawner.NewMoveOneVelocitySpawner(startVelocity, len(sides))
	posGen := position_spawner.NewBoundedGenerator(sides)

	containerInst := container.NewRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()

	spawnerInst := spawner.NewSpawnerImpl(radius, mass, containerInst, posGen, velSpawner)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

	visualisation := visualisation.StartVisualisation(dim)

  visualisation.Init(simulationInst.Particles())

	simulationInst.Observers().SubscribeParticle(visualisation)
	simulationInst.Observers().SubscribeCollision(visualisation)
	simulationInst.Observers().SubscribeCollisionWithContainer(visualisation)

	defer visualisation.Report()

	simulationInst.Run(-1.0)
}
