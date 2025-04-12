package naive_chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type NaiveChunkFactory struct {
	dt float64
}

func NewNaiveChunkFactory(dt float64) chunk.ChunkFactory {
	return &NaiveChunkFactory{dt: dt}
}

func (f NaiveChunkFactory) NewChunk(c container.Container, p []particle.Particle) chunk.Chunk {
	return newNaiveChunk(f.dt, c, p)
}
