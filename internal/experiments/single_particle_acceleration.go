package experiments

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	position_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
)

func RunSingleParticleAcceleration(dim int) {
	// argon
	radius := 1.88e-10
	mass := 6.6335e-26

	side := 80e-9
	count := 10000
	maxVelocity := 400.0

	sides := make([]float64, dim)
	for i := range sides {
		sides[i] = side
	}

	containerInst := container.NewSimpleRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()

	velSpawner := velocity_spawner.NewMoveOneVelocitySpawner(maxVelocity*1000, len(sides))
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
	fmt.Println("Simulation starting")

	simulationInst.Run(5e-13)
}
