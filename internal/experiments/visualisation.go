package experiments

import (
	// "github.com/iv4n-t3a/fart-simulator/internal/chunk/kdtree_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/chunk/naive_chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/simulation"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner"
	"github.com/iv4n-t3a/fart-simulator/internal/visualisation"
)

func RunVisualisation() {
	side := 0.05
  dt := 1e-7
  max_velocity := 1.0
  radius := 0.0001
  mass := 1.0
  count := 100

	sides := []float64{side, side, side}
	containerInst := container.NewRectContainer(sides)
	chunkFactory := naive_chunk.NewNaiveChunkFactory(dt)
	// chunkFactory := kdtree_chunk.NewKDTreeChunkFactory()
	spawnerInst := spawner.NewRectSpawner(max_velocity, radius, mass, *containerInst)

	simulationInst := simulation.NewSingleChunkSimulation(count, containerInst, chunkFactory, spawnerInst)

  visualisation := visualisation.StartVisualisation()
  simulationInst.Observers().SubscribeParticle(visualisation)
  defer visualisation.Report()

	simulationInst.Run(-1.0)
}
