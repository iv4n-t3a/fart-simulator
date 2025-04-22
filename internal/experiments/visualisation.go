package experiments

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunVisualisation() {
	side := 0.05

	sides := []float64{side, side, side}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()
	spawnerInst := spawner.NewRectSpawner(1.0, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(10000, containerInst, chunkFactory, spawnerInst)

  visualisation := visualisation.StartVisualisation()
  simulationInst.Observers().SubscribeParticle(visualisation)
  defer visualisation.Report()

	simulationInst.Run(-1.0)
}
