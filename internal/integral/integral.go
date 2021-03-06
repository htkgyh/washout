// Copyright © 2017 shoarai

package integral

type Integral struct {
	interval uint
	sums     []float64
}

func New(interval uint) *Integral {
	return NewMulti(interval, 1)
}

func NewMulti(interval uint, number uint) *Integral {
	return &Integral{interval, make([]float64, number)}
}

func (integ *Integral) Integrate(val float64) float64 {
	addedVal := val
	for i := range integ.sums {
		integ.sums[i] += addedVal * float64(integ.interval) / 1000
		addedVal = integ.sums[i]
	}
	return integ.sums[len(integ.sums)-1]
}
