package main

import (
	"books/chapter11/math"
	"fmt"
)

func main() {
	xs := []float64{-5,1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)

	e := []float64{}
	fmt.Println(math.Average(e))

	fmt.Println(math.Min(xs))
	fmt.Println(math.Min(e))

	fmt.Println(math.Max(xs))
	fmt.Println(math.Max(e))

}
