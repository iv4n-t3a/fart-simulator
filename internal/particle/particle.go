package particle

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/config"
	"github.com/iv4n-t3a/fart-simulator/internal/math_util"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type Particle struct {
	Pos    vector.Vector
	Vel    vector.Vector
	Radius float64
	Mass   float64
}

func ProcessCollision(a *Particle, b *Particle) bool {
	if a.Pos.Dist(b.Pos) > a.Radius+b.Radius {
		return false
	}

	a.Vel = velocityAfterCollision(*a, *b)
	b.Vel = velocityAfterCollision(*b, *a)
	return true
}

func velocityAfterCollision(a Particle, b Particle) vector.Vector {
	d := a.Pos.Sub(b.Pos).Normalized()
	v1 := a.Vel.DotProd(d)
	v2 := b.Vel.DotProd(d)
	sumImp := v1*a.Mass + v2*b.Mass
	sumEnergy := v1*v1*a.Mass/2 + v2*v2*b.Mass/2

	// Solves the following system of equations
	// P = v1 * m1 + v2 * m2
	// E = v1^2 * m1 / 2 + v2^2 * m2 / 2
	x1, x2, ok := math_util.SolveSqrEquation(
		b.Mass+b.Mass*b.Mass/a.Mass/a.Mass,
		2*b.Mass*sumImp/a.Mass,
		sumImp*sumImp/a.Mass/a.Mass-2*sumEnergy,
	)

	if !ok {
		panic("Equivalation is supposed to have solution")
	}

  new_v1 := x1
  if math.Abs(x1 - v1) <= config.Eps {
    new_v1 = x2
  }

	return a.Vel.Sub(d.Mul(v1)).Add(d.Mul(new_v1))
}
