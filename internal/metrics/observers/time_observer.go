package observers

type TimeObserver struct {
	Duration float64
}

func NewTimeObserver() *TimeObserver {
	return &TimeObserver{}
}

func (t *TimeObserver) Tick(dt float64) {
	t.Duration += dt
}
