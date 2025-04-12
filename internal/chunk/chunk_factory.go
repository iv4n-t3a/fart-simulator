package chunk

import (
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
)

type ChunkFactory interface {
	NewChunk(container.Container, []particle.Particle) Chunk
}
