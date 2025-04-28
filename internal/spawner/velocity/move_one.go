package velocity_spawner

import "github.com/iv4n-t3a/fart-simulator/internal/vector"

type MoveOneVelocitySpawner struct {
	innerSpawner VelocitySpawner
	dim          int
	moved        bool
}

func NewMoveOneVelocitySpawner(vel float64, dim int) VelocitySpawner {
	return &MoveOneVelocitySpawner{
		innerSpawner: NewSameVelocitySpawner(vel, dim),
		dim:          dim,
		moved:        false,
	}
}

func (s *MoveOneVelocitySpawner) NextVelocity() vector.Vector {
	if s.moved {
		return vector.ZeroVector(s.dim)
	} else {
		s.moved = true
		return s.innerSpawner.NextVelocity()
	}
}
