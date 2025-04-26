package position_spawner

import "github.com/iv4n-t3a/fart-simulator/internal/vector"

type PositionGenerator interface {
	NextPosition() vector.Vector
}
