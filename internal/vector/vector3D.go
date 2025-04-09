package vector

import (
	"fmt"
	"math"
)

type Vector3D struct {
	x float64
	y float64
	z float64
}

func (v Vector3D) Add(oth Vector) Vector {
	return Vector3D{
		x: v.x + oth.X(),
		y: v.y + oth.Y(),
		z: v.z + oth.Z(),
	}
}

func (v Vector3D) Sub(oth Vector) Vector {
	return Vector3D{
		x: v.x - oth.X(),
		y: v.y - oth.Y(),
		z: v.z - oth.Z(),
	}
}

func (v Vector3D) Mul(num float64) Vector {
	return Vector3D{
		x: v.x * num,
		y: v.y * num,
		z: v.z * num,
	}
}

func (v Vector3D) Div(num float64) Vector {
	if num == 0.0 {
		panic("Dividing vector by zero")
	}
	return Vector3D{
		x: v.x / num,
		y: v.y / num,
		z: v.z / num,
	}
}

func (v Vector3D) DotProd(oth Vector) float64 {
	return v.x*oth.X() + v.y*oth.Y() + v.z*oth.Z()
}

func (v Vector3D) CrossProd(oth Vector) Vector {
	panic("Unimplemented")
}

func (v Vector3D) Normalized() Vector {
	return v.Div(v.Lenght())
}

func (v Vector3D) Lenght() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
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
	panic(fmt.Sprintf("Trying to access %d dimension on 2D vector", i))
}
