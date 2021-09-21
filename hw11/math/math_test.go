package math

import (
	"math"
	"testing"
)

type testpair struct {
	values []float64
	average float64
	min float64
	max float64
}

var tests = []testpair{
	{ []float64{}, 0, math.NaN(), math.NaN()},
	{ []float64{1,2}, 1.5, 1,2 },
	{ []float64{1,1,1,1,1,1}, 1, 1,1 },
	{ []float64{-1,1}, 0,-1,1 },
	{ []float64{1,2}, 1.5, 1,2 },
}
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
func TestMin(t *testing.T) {
	for _, pair := range tests {
		v := Min(pair.values)
		if len(pair.values) == 0 && math.IsNaN(v) {
			continue
		}
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
func TestMax(t *testing.T) {
	for _, pair := range tests {
		v := Max(pair.values)
		if len(pair.values) == 0 && math.IsNaN(v) {
			continue
		}
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}