package main

import (
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

func RunSimpleSimulation() {
	side := config.ContainerSide
	time := config.SimulationTime

	sides := []float64{side, side, side}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()
	spawnerInst := spawner.NewRectSpawner(1, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(1000, containerInst, chunkFactory, spawnerInst)

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

func main() {
	RunSimpleSimulation()
}
