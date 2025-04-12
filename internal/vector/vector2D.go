package vector

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/config"
	"math"
)

type Vector2D struct {
	x float64
	y float64
}

func NewVector2D(x float64, y float64) Vector2D {
	return Vector2D{
		x: x,
		y: y,
	}
}

func (v Vector2D) Add(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return &Vector2D{
		x: v.x + oth.X(),
		y: v.y + oth.Y(),
	}
}

func (v Vector2D) Sub(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return &Vector2D{
		x: v.x - oth.X(),
		y: v.y - oth.Y(),
	}
}

func (v Vector2D) Mul(num float64) Vector {
	return &Vector2D{
		x: v.x * num,
		y: v.y * num,
	}
}

func (v Vector2D) Div(num float64) Vector {
	if num == 0.0 {
		panic("Dividing vector by zero")
	}
	return &Vector2D{
		x: v.x / num,
		y: v.y / num,
	}
}

func (v Vector2D) DotProd(oth Vector) float64 {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return v.x*oth.X() + v.y*oth.Y()
}

func (v Vector2D) CrossProd(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	panic("Cross product is not supported for 2D")
}

func (v Vector2D) Normalized() Vector {
	if v.Length() == 0 {
		return &Vector2D{0, 0}
	}
	return v.Div(v.Length())
}

func (v Vector2D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v Vector2D) Dist(oth Vector) float64 {
	return v.Sub(oth).Length()
}

func (v Vector2D) IsCollinear(oth Vector) bool {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	vNorm := v.Normalized()
	othNorm := oth.Normalized()
	return vNorm.Sub(othNorm).Length() < config.Eps || vNorm.Add(othNorm).Length() < config.Eps
}

func (v Vector2D) X() float64 {
	return v.x
}

func (v Vector2D) Y() float64 {
	return v.y
}

func (v Vector2D) Z() float64 {
	panic("Trying to access Z in 2D vector")
}

func (v Vector2D) Dimensions() int {
	return 2
}

func (v Vector2D) Dimension(i int) float64 {
	if i == 0 {
		return v.x
	}
	if i == 1 {
		return v.y
	}
	panic(fmt.Sprintf("Trying to access %d dimension on 2D vector", i))
}

func (v Vector2D) SetDimension(val float64, i int) {
	if i == 0 {
		v.x = val
		return
	}
	if i == 1 {
		v.y = val
		return
	}
	panic(fmt.Sprintf("Trying to access %d dimension on 2D vector", i))
}
