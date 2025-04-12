package vector

import (
	"fmt"
	"github.com/iv4n-t3a/fart-simulator/config"
	"math"
)

type Vector3D struct {
	x float64
	y float64
	z float64
}

func (v Vector3D) Coords() []float64 {
	return []float64{v.x, v.y, v.z}
}

func NewVector3D(x float64, y float64, z float64) Vector3D {
	return Vector3D{
		x: x,
		y: y,
		z: z,
	}
}

func (v Vector3D) Add(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return &Vector3D{
		x: v.x + oth.X(),
		y: v.y + oth.Y(),
		z: v.z + oth.Z(),
	}
}

func (v Vector3D) Sub(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return &Vector3D{
		x: v.x - oth.X(),
		y: v.y - oth.Y(),
		z: v.z - oth.Z(),
	}
}

func (v Vector3D) Mul(num float64) Vector {
	return &Vector3D{
		x: v.x * num,
		y: v.y * num,
		z: v.z * num,
	}
}

func (v Vector3D) Div(num float64) Vector {
	if num == 0.0 {
		panic("Dividing vector by zero")
	}
	return &Vector3D{
		x: v.x / num,
		y: v.y / num,
		z: v.z / num,
	}
}

func (v Vector3D) DotProd(oth Vector) float64 {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return v.x*oth.X() + v.y*oth.Y() + v.z*oth.Z()
}

func (v Vector3D) CrossProd(oth Vector) Vector {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	return &Vector3D{
		x: v.y*oth.Z() - v.z*oth.Y(),
		y: v.z*oth.X() - v.x*oth.Z(),
		z: v.x*oth.Y() - v.y*oth.X(),
	}
}

func (v Vector3D) Normalized() Vector {
	if v.Length() == 0 {
		return &Vector3D{0, 0, 0}
	}
	return v.Div(v.Length())
}

func (v Vector3D) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector3D) Dist(oth Vector) float64 {
	return v.Sub(oth).Length()
}

func (v Vector3D) IsCollinear(oth Vector) bool {
	if v.Dimensions() != oth.Dimensions() {
		panic(fmt.Sprintf("Dimensions %d != %d", v.Dimensions(), oth.Dimensions()))
	}
	vNorm := v.Normalized()
	othNorm := oth.Normalized()
	return vNorm.Sub(othNorm).Length() < config.Eps || vNorm.Add(othNorm).Length() < config.Eps
}

func (v Vector3D) X() float64 {
	return v.x
}

func (v Vector3D) Y() float64 {
	return v.y
}

func (v Vector3D) Z() float64 {
	return v.z
}

func (v Vector3D) Dimensions() int {
	return 3
}

func (v Vector3D) Dimension(i int) float64 {
	if i == 0 {
		return v.x
	}
	if i == 1 {
		return v.y
	}
	if i == 2 {
		return v.z
	}
	panic(fmt.Sprintf("Trying to access %d dimension on 3D vector", i))
}

func (v Vector3D) SetDimension(val float64, i int) Vector {
	switch i {
	case 0:
		v.x = val
	case 1:
		v.y = val
	case 2:
		v.z = val
	default:
		panic(fmt.Sprintf("Trying to access %d dimension on 3D vector", i))
	}
	return v
}
