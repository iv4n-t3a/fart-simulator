package spawner

import "github.com/iv4n-t3a/fart-simulator/internal/particle"

type Spawner interface {
	SpawnParticle() particle.Particle
}
