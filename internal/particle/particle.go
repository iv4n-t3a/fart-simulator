package particle

import (
	"fmt"
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

func (p Particle) Impulse() vector.Vector {
	return p.Vel.Mul(p.Mass)
}

func (p Particle) KineticEnergy() float64 {
	return p.Vel.Length() * p.Vel.Length() * p.Mass / 2
}

func ProcessCollision(p1 *Particle, p2 *Particle) bool {
	if p1.Pos.Dist(p2.Pos) > p1.Radius+p2.Radius {
		return false
	}

	collisionDir := p2.Pos.Sub(p1.Pos).Normalized()
	w1 := collisionDir.Mul(p1.Vel.DotProd(collisionDir))
	w2 := collisionDir.Mul(p2.Vel.DotProd(collisionDir))
	u1 := p1.Vel.Sub(w1)
	u2 := p2.Vel.Sub(w2)

	// To avoid solving square equation in case of same mass
	if math.Abs(p1.Mass-p2.Mass) < config.Eps {
		p1.Vel = u1.Add(w2)
		p2.Vel = u2.Add(w1)
		return true
	}

	new_w1, new_w2 := velocitiesAfterCollision(w1.Length(), w2.Length(), p1.Mass, p2.Mass)
	p1.Vel = u1.Add(collisionDir.Mul(new_w1))
	p2.Vel = u2.Add(collisionDir.Mul(new_w2))
	return true
}

func velocitiesAfterCollision(u1 float64, u2 float64, m1 float64, m2 float64) (float64, float64) {
	P := u1*m1 + u2*m2
	E := u1*u1*m1/2 + u2*u2*m2/2

	// Solves the following system of equations
	// m1 u1 + m2 u2 = P
	// m1 u1**2 + m2 u2**2 = 2E
	u11, u12, ok := math_util.SolveSqrEquation(m1*(m1+m2), -2*m1*P, P*P-2*E*m2)

	if !ok {
		panic("Equation is supposed to have solutions")
	}

	if math.Abs(u1-u11) <= config.Eps {
		u1 = u12
	} else if math.Abs(u1-u12) <= config.Eps {
		u1 = u11
	} else {
		panic(fmt.Sprintf("One solution is supposed to match with input: %f != %f, %f", u1, u11, u12))
	}

	u2 = (P - m1*u1) / m2

	return u1, u2
}
