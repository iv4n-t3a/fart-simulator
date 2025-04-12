package vector

import "fmt"

type VectorImpl struct {
	coordinates []float64
}

func NewVector(coordinates []float64) Vector {
	switch len(coordinates) {
	case 2:
		return NewVector2D(coordinates[0], coordinates[1])
	case 3:
		return NewVector3D(coordinates[0], coordinates[1], coordinates[2])
	default:
		// return VectorImpl{coordinates: coordinates}
		panic(fmt.Sprintf("Vector type is not implemented for %d dimensions", len(coordinates)))
	}
}

func ZeroVector(dim int) Vector {
	switch dim {
	case 2:
		return NewVector2D(0, 0)
	case 3:
		return NewVector3D(0, 0, 0)
	default:
		// return VectorImpl{coordinates: make([]float64, dim)}
		panic(fmt.Sprintf("Vector type is not implemented for %d dimensions", dim))
	}
}

// Not implemented
