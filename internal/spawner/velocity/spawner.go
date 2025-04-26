package velocity_spawner

import "github.com/iv4n-t3a/fart-simulator/internal/vector"

type VelocitySpawner interface {
  NextVelocity() vector.Vector
}
