package position_spawner

import (
	// "github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"github.com/kyroy/kdtree"
)

type PositionSpawner struct {
	created   kdtree.KDTree
	generator PositionGenerator
	container container.Container
	count     int
}

type item struct {
	vector.Vector
	radius float64
}

func NewPositionSpawner(generator PositionGenerator, container container.Container) *PositionSpawner {
	return &PositionSpawner{
		generator: generator,
		container: container,
		count:     0,
	}
}

func (s *PositionSpawner) NextPosition(radius float64) vector.Vector {
	p := s.generator.NextPosition()
	for !s.canInsert(p, radius) {
	}
	s.created.Insert(
		item{
			Vector: p,
			radius: radius,
		})

	s.count += 1
	return p
}

func (s *PositionSpawner) canInsert(pos vector.Vector, radius float64) bool {
	if !s.container.IsInside(pos) {
		return false
	}
	if s.count <= 1 {
		return true
	}
	nearest := s.created.KNN(pos, 2)[1].(item)
	return pos.Dist(nearest) >= nearest.radius+radius+config.Eps
}
