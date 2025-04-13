package main

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

func RunSimpleSimulation() {
  side := 1e-5
  time := 1e-7

	sides := []float64{side, side, side}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()
	spawnerInst := spawner.NewRectSpawner(1.0, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(10000, containerInst, chunkFactory, spawnerInst)

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

	simulationInst.Run(time)
}

func main() {
	RunSimpleSimulation()
}
