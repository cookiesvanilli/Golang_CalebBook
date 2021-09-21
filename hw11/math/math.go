package math

import "math"

func Average(xs []float64) float64 {
	if len(xs) != 0 {
		total := float64(0)
		for _, x := range xs {
			total += x
		}
		return total / float64(len(xs))
	} else {
		return 0
	}
}
func Min(list []float64) float64 {
	if len(list) == 0 {
		return math.NaN()
	}
	if len(list) == 0 {
		return 0
	}
	min := list[0]
	for _, x := range list {
		if min > x {
			min = x
		}
	}
	return min
}
func Max(list []float64) float64  {
	if len(list) == 0 {
		return math.NaN()
	}
	if len(list) == 0 {
		return 0
	}
	max := list[0]
	for _, x := range list {
		if max < x {
			max = x
		}
	}
	return max
}
