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

func RunShrinkingContainerFile(dim int) {
	// argon
	radius := 1.88e-10
	mass := 6.6335e-26

	side := 40e-8
	count := 10000
	maxVelocity := 400.0
	shrinkingSpeed := 100.0
	shrinkingResistance := mass * 50

	sides := make([]float64, dim)
	for i := range sides {
		sides[i] = side
	}

	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()
	containerInst := container.NewShrinkingRectContainer(sides, shrinkingSpeed, shrinkingResistance)
	velSpawner := velocity_spawner.NewNaiveVelocitySpawner(maxVelocity, len(sides))
	posGen := position_spawner.NewBoundedGenerator(sides)
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

	particleObserver := metrics.NewParticleObserver(timeObserver)
	simulationInst.Observers().SubscribeParticle(particleObserver)
	defer particleObserver.Report()

	simulationInst.Run(10e-13)
}
