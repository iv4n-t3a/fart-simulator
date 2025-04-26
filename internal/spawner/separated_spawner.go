package spawner

import (
	"github.com/iv4n-t3a/fart-simulator/internal/container"
	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/spawner/position"
	velocity_spawner "github.com/iv4n-t3a/fart-simulator/internal/spawner/velocity"
	"math/rand"
)

type SeparatedPositionSpawner struct {
	pos1                  *position_spawner.PositionSpawner
	pos2                  *position_spawner.PositionSpawner
	vel1                  velocity_spawner.VelocitySpawner
	vel2                  velocity_spawner.VelocitySpawner
	container             container.RectContainer
	spaceProportion       float64
	probabilityProportion float64
	mass1                 float64
	mass2                 float64
	radius1               float64
	radius2               float64
	nextIndex             int64
}

func NewSeparatedPositionSpawner(
	vel1 velocity_spawner.VelocitySpawner,
	vel2 velocity_spawner.VelocitySpawner,
	spaceProportion float64,
	probabilityProportion float64,
	containerInst container.RectContainer,
	mass1 float64,
	mass2 float64,
	radius1 float64,
	radius2 float64) *SeparatedPositionSpawner {
	sides := containerInst.GetSides()

	sides1 := append([]float64(nil), sides...)
	sides1[0] *= spaceProportion
	container1 := container.NewSimpleRectContainer(sides1)
	generator1 := position_spawner.NewBoundedGenerator(sides1)
	spawner1 := position_spawner.NewPositionSpawner(generator1, container1)

	sides2 := append([]float64(nil), sides...)
	sides2[0] *= 1 - spaceProportion
	container2 := container.NewSimpleRectContainer(sides2)
	generator2 := position_spawner.NewBoundedGenerator(sides2)
	spawner2 := position_spawner.NewPositionSpawner(generator2, container2)

	return &SeparatedPositionSpawner{
		pos1:                  spawner1,
		pos2:                  spawner2,
		vel1:                  vel1,
		vel2:                  vel2,
		container:             containerInst,
		spaceProportion:       spaceProportion,
		probabilityProportion: probabilityProportion,
		mass1:                 mass1,
		mass2:                 mass2,
		radius1:               radius1,
		radius2:               radius2,
	}
}

func (s *SeparatedPositionSpawner) SpawnParticle() particle.Particle {
	decider := rand.Float64()
	s.nextIndex++
	if decider <= s.probabilityProportion {
		return particle.Particle{
			Pos:    s.pos1.NextPosition(s.radius1),
			Vel:    s.vel1.NextVelocity(),
			Radius: s.radius1,
			Mass:   s.mass1,
			Index:  s.nextIndex - 1,
		}
	} else {
		newPos := s.pos2.NextPosition(s.radius2)
		newPos = newPos.SetDimension(newPos.Dimension(0)+s.container.GetSides()[0]*s.spaceProportion, 0)
		return particle.Particle{
			Pos:    newPos,
			Vel:    s.vel2.NextVelocity(),
			Radius: s.radius2,
			Mass:   s.mass2,
			Index:  s.nextIndex - 1,
		}
	}
}
