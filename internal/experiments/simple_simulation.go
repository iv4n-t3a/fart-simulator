package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
)

func RunSimpleSimulation(dim int) {
	side := 0.05
	time := 1.0
	count := 10000
  radius := 1e-4
  mass := 1.5e-20

	sides := make([]float64, dim)

	for i := range sides {
		sides[i] = side
	}

	containerInst := container.NewSimpleRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()


	velSpawner := velocity_spawner.NewNaiveVelocitySpawner(1.0, len(sides))
	posGen := position_spawner.NewBoundedGenerator(sides)

	spawnerInst := spawner.NewSpawnerImpl(radius, mass, containerInst, posGen, velSpawner)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

	timeObserver := metrics.NewTimeObserver()
	simulationInst.Observers().SubscribeTime(timeObserver)

	colCounter := metrics.NewCollisionCounterObserver()
	simulationInst.Observers().SubscribeCollision(colCounter)
	defer colCounter.Report()

	colContCounter := metrics.NewCollisionWithContainerCounterObserver()
	simulationInst.Observers().SubscribeCollisionWithContainer(colContCounter)
	defer colContCounter.Report()

	containerAggregator := metrics.NewCollisionWithContainerAggregatorObserver(timeObserver)
	simulationInst.Observers().SubscribeCollisionWithContainer(containerAggregator)
	defer containerAggregator.Report()

	particleObserver := metrics.NewParticleObserver(timeObserver)
	simulationInst.Observers().SubscribeParticle(particleObserver)
	defer particleObserver.Report()

	simulationInst.Run(time)
}
