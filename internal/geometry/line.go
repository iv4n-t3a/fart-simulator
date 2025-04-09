package geometry

import (
	"fmt"
	"math"

	"github.com/iv4n-t3a/fart-simulator/internal/vector"
)

type Line struct {
	Start vector.Vector
	Dir   vector.Vector
}

func NewLine(start vector.Vector, dir vector.Vector) Line {
  if start.Dimensions() != dir.Dimensions() {
    panic(fmt.Sprintf("Dimensions mismatch %d, %d", start.Dimensions(), dir.Dimensions()))
  }
  if dir.Length() == 0.0 {
    panic("Null dir vector")
  }
  return Line{Start: start, Dir: dir}
}

func (l Line) Dimensions() int {
  return l.Dir.Dimensions()
}

func (l1 Line) Dist(l2 Line) float64 {
  if l1.Dimensions() != l2.Dimensions() {
    panic(fmt.Sprintf("Dimensions mismatch %d, %d", l1.Dimensions(), l2.Dimensions()))
  }
  d := l1.Dir.CrossProd(l2.Dir)
  d = d.Normalized()
  return math.Abs(d.DotProd(l1.Start) - d.DotProd(l2.Start))
}
