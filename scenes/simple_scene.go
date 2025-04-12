package scenes

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/observers_impl"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

func RunSimpleScene() {
	rectContainerInst := container.NewRectContainer([]float64{1, 1})
	naiveChunk := chunk.NewNaiveChunk(1.0/100, rectContainerInst)
	particles := particle.SpawnStaticRectangular2D(
		100,
		vector.NewVector2D(0, 0),
		vector.NewVector2D(1, 1),
	)
	(*particles)[0].Vel = vector.NewVector2D(10, 10)

	naiveChunk.InitializeParticles(particles)

	colCounter := observers_impl.NewCollisionCounterObserver()
	colContCounter := observers_impl.NewCollisionWithContainerCounterObserver()

	naiveChunk.SubscribeCollision(colCounter)
	naiveChunk.SubscribeCollisionWithContainer(colContCounter)

	simulationInst := simulation.NewSingleChunkSimulation(naiveChunk, nil)
	simulationInst.AddReporter(colCounter)
	simulationInst.AddReporter(colContCounter)

	simulationInst.Run(1)
	simulationInst.ReportMetrics()
}
