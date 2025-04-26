package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunWallWithHole(dim int) {
	side := 0.05
	dt := 1e-6
	startVelocity := 1.0
	count := 100
	radius := 1e-4
	mass := 1.5e-20
	holeRadius := radius * 10

	sides := make([]float64, dim)
	holePos := make([]float64, dim)
  spawnArea := make([]float64, dim)

	for i := range sides {
		sides[i] = side
		holePos[i] = side / 2
    spawnArea[i] = side
	}
  spawnArea[0] = side / 2

	velSpawner := velocity_spawner.NewNaiveVelocitySpawner(startVelocity, len(sides))
	posGen := position_spawner.NewBoundedGenerator(spawnArea)

	wall := container.NewWallWithHole(vector.NewVector(holePos), 0, holeRadius)
	externalContainer := container.NewSimpleRectContainer(sides)
	containerInst := container.Pair(wall, externalContainer)

	chunkFactory := naive_chunk.NewNaiveChunkFactory(dt)

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
