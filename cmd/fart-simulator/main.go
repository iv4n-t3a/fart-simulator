package main

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	simulation2 "github.com/iv4n-t3a/fart-simulator/internal/simulation"
)

func main() {
	naiveChunk := chunk.NewNaiveChunk(1.0 / 100)
	simulation := simulation2.NewSingleChunkSimulation(naiveChunk, nil)
	simulation.Run(1)
}
