package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

func RunHole(dim int) {
	// argon
	radius := 1.88e-10
	mass := 6.6335e-26

	side := 80e-9
	count := 1000
	startVelocity := 400.0
	duration := 5e-12

	holeRadius := side * 0.1

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

	chunkFactory := naive_chunk.NewNaiveChunkFactory(duration / 50000)

	spawnerInst := spawner.NewSpawnerImpl(radius, mass, containerInst, posGen, velSpawner)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

	colCounter := metrics.NewCollisionCounterObserver()
	simulationInst.Observers().SubscribeCollision(colCounter)
	defer colCounter.Report()

	timeObserver := metrics.NewTimeObserver()
	simulationInst.Observers().SubscribeTime(timeObserver)

	containerAggregator := metrics.NewCollisionWithContainerAggregatorObserver(timeObserver)
	simulationInst.Observers().SubscribeCollisionWithContainer(containerAggregator)
	defer containerAggregator.Report()

	simulationInst.Run(duration)
}
