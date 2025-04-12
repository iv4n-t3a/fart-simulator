package main

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

func RunSimpleSimulation() {
	sides := []float64{0.002, 0.002, 0.002}
	container := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(0.01)
	spawner := spawner.NewRectSpawner(5.0, *container)

	simulationInst := simulation.NewSingleChunkSimulation(1000, container, chunkFactory, spawner)

	colCounter := observers.NewCollisionCounterObserver()
	simulationInst.SubscribeCollision(colCounter)
	simulationInst.AddReporter(colCounter)

	colContCounter := observers.NewCollisionWithContainerCounterObserver()
	simulationInst.SubscribeCollisionWithContainer(colContCounter)
	simulationInst.AddReporter(colContCounter)

	simulationInst.Run(1.0)
	simulationInst.ReportMetrics()
}

func main() {
	RunSimpleSimulation()
}
