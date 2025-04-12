package chunk

import "github.com/iv4n-t3a/fart-simulator/internal/container"

type ChunkFactory interface {
	NewChunk(container.Container) Chunk
}
