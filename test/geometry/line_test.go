package geometry_test

import (
	"testing"

	"github.com/iv4n-t3a/fart-simulator/internal/geometry"
	"github.com/iv4n-t3a/fart-simulator/internal/vector"
	"github.com/stretchr/testify/assert"
)

func TestDist(t *testing.T) {
	l1 := geometry.NewLine(
		vector.NewVector3D(0.0, 0.0, 0.0),
		vector.NewVector3D(0.0, 1.0, 1.0),
	)
	l2 := geometry.NewLine(
		vector.NewVector3D(1.0, 0.0, 0.0),
		vector.NewVector3D(0.0, -1.0, 1.0),
	)

	dist := l1.Dist(l2)
	assert.InEpsilon(t, dist, 1.0, 0.001, "Dist")
}

func TestDistParallel(t *testing.T) {
	l1 := geometry.NewLine(
		vector.NewVector3D(0.0, 0.0, 0.0),
		vector.NewVector3D(0.0, 0.0, 1.0),
	)
	l2 := geometry.NewLine(
		vector.NewVector3D(1.0, 0.0, 0.0),
		vector.NewVector3D(0.0, 0.0, 1.0),
	)

	dist := l1.Dist(l2)
	assert.InEpsilon(t, dist, 1.0, 0.001, "Dist")
}
