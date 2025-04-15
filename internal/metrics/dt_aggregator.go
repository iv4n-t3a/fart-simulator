package metrics

import (
	"fmt"
	"math"
)

type DtAggregator struct {
	meanDt       float64
	meanSquareDt float64
	minDt        float64
	maxDt        float64

	measurements int64
}

func NewDtAggregator() *DtAggregator {
	return &DtAggregator{
		meanDt:       0.0,
		meanSquareDt: 0.0,
		minDt:        math.Inf(1),
		maxDt:        0.0,
		measurements: 0,
	}
}

func (t *DtAggregator) Tick(dt float64) {
	t.minDt = min(dt, t.minDt)
	t.maxDt = max(dt, t.maxDt)
	if t.measurements == 0 {
		t.measurements = 1
		t.meanDt = dt
		t.meanSquareDt = dt
		return
	}
	t.meanDt = (t.meanDt*float64(t.measurements) + dt) / float64(t.measurements+1)
	t.meanSquareDt = math.Sqrt((t.meanSquareDt*t.meanSquareDt*float64(t.measurements) + dt*dt) / float64(t.measurements+1))
	t.measurements += 1
}

func (t *DtAggregator) Report() {
	fmt.Println("Mean dt:", t.meanDt)
	fmt.Println("Mean square dt:", t.meanSquareDt)
	fmt.Println("Min dt:", t.minDt)
	fmt.Println("Max dt:", t.maxDt)
}
