package kdtree_chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/chunk"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type KDTreeChunkFactory struct {
}

func NewKDTreeChunkFactory() chunk.ChunkFactory {
	return &KDTreeChunkFactory{}
}

func (f KDTreeChunkFactory) NewChunk(c container.Container, p []particle.Particle) chunk.Chunk {
	return newKDTreeChunk(c, p)
}
