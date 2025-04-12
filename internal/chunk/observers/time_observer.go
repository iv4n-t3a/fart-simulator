package observers

type TimeObserver interface {
	Tick(float64)
}
