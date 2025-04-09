package vector

type Vector interface {
	Add(Vector) Vector
	Sub(Vector) Vector
	Mul(float64) Vector
	Div(float64) Vector
	DotProd(Vector) float64
	CrossProd(Vector) Vector
	Normalized() Vector
	Length() float64
  Dist(Vector) float64
	IsCollinear(Vector) bool

	X() float64
	Y() float64
	Z() float64 // Will panic on Vector2D

	// kdtree.Point compatability
	Dimensions() int
	Dimension(i int) float64
}
