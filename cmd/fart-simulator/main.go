package main

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

func RunSimpleSimulation() {
	sides := []float64{0.05, 0.05, 0.05}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(0.0001)
	spawnerInst := spawner.NewRectSpawner(1.0, *containerInst)

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

	simulationInst.Run(1.0)
}

func main() {
	RunSimpleSimulation()
}
