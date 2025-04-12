package main

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/metrics/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
)

func RunSimpleSimulation() {
	sides := []float64{1.0, 1.0, 1.0}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(0.01)
	spawnerInst := spawner.NewRectSpawner(5.0, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(100, containerInst, chunkFactory, spawnerInst)

	colCounter := observers.NewCollisionCounterObserver()
	simulationInst.SubscribeCollision(colCounter)
	simulationInst.AddReporter(colCounter)

	colContCounter := observers.NewCollisionWithContainerCounterObserver()
	simulationInst.SubscribeCollisionWithContainer(colContCounter)
	simulationInst.AddReporter(colContCounter)

	timeObserver := observers.NewTimeObserver()
	simulationInst.SubscribeTime(timeObserver)
	ContainerPlot := observers.NewCollisionWithContainerPlotObserver(timeObserver)
	simulationInst.SubscribeCollisionWithContainer(ContainerPlot)
	simulationInst.AddReporter(ContainerPlot)

	simulationInst.Run(1.0)
	simulationInst.ReportMetrics()
}

func main() {
	RunSimpleSimulation()
}
