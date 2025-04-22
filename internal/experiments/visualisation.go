package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunVisualisation() {
	side := 0.05

	sides := []float64{side, side, side}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(0.0000001)
	spawnerInst := spawner.NewRectSpawner(1.0, 0.00001, 1.0, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(100, containerInst, chunkFactory, spawnerInst)

  visualisation := visualisation.StartVisualisation()
  simulationInst.Observers().SubscribeParticle(visualisation)
  defer visualisation.Report()

	simulationInst.Run(-1.0)
}
