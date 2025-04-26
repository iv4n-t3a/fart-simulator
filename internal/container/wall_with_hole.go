package container

import (
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/particle"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type WallWithHole struct {
	// Wall position is set by dimension and position in it
	dim int
	pos float64

	holeRadius   float64
	holePosition vector.Vector // position.Dimension(dim) == pos
}

func NewWallWithHole(holePosition vector.Vector, wallDimension int, holeRadius float64) *WallWithHole {
	return &WallWithHole{
		dim:          wallDimension,
		pos:          holePosition.Dimension(wallDimension),
		holeRadius:   holeRadius,
		holePosition: holePosition,
	}
}

func (c *WallWithHole) IsInside(v vector.Vector) bool {
	// infinite container
	return true
}

func (c *WallWithHole) ProcessCollision(p *particle.Particle) bool {
	if c.holeRadius >= p.Radius && p.Pos.Dist(c.holePosition) <= c.holeRadius-p.Radius {
		return false
	}

	if math.Abs(p.Dimension(c.dim) - c.pos) <= p.Radius {
		p.Vel = p.Vel.SetDimension(-p.Vel.Dimension(c.dim), c.dim)
		return true
	}
	return false
}

func (c *WallWithHole) TimeBeforeCollision(p particle.Particle) float64 {
	t := math.Inf(1)
	v := p.Vel.Dimension(c.dim)
	x := p.Pos.Dimension(c.dim)

	if v > 0 {
		t = min((c.pos-x-p.Radius)/v, t)
	} else {
		t = min((p.Radius-x)/v, t)
	}

  nextPos := p.Pos.Add(p.Vel.Mul(t))

  if c.holeRadius >= p.Radius && nextPos.Dist(c.holePosition) <= c.holeRadius-p.Radius {
		return math.Inf(1)
	}

	return t
}
