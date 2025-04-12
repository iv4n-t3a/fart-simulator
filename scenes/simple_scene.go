package scenes

import (
	"github.com/iv4n-t3a/fart-simulator/internal/metrics/observers"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
)

func RunSimpleScene() {
	colCounter := observers.NewCollisionCounterObserver()
	colContCounter := observers.NewCollisionWithContainerCounterObserver()

	simulationInst := simulation.NewSingleChunkSimulation(100)
	simulationInst.SubscribeCollision(colCounter)
	simulationInst.SubscribeCollisionWithContainer(colContCounter)
	simulationInst.AddReporter(colCounter)
	simulationInst.AddReporter(colContCounter)

	simulationInst.Run(1)
	simulationInst.ReportMetrics()
}
