package main

import "fmt"

func main() {
	averageTotal()
	ff()
	returned()
	close()
	close2()
	generator()
	recursion()
	panick()
	panick2()
}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func averageTotal() {
	someOtherName := []float64{98,93,77,82,83}
	fmt.Println(average(someOtherName))
}

func ff() {
	fmt.Println(f1())
}
func f1() int {
	return f2()
}
func f2() int {
	return 1
}

func f() (x int, err int) {
	return 5, 6
}
func returned() {
	x, err := f()
	fmt.Println(x, err)
}
// Замыкания ------------------------->
func close() {
	add := func(x, y int) int {
		return x+y
	}
	fmt.Println(add(1,1))
}
func close2() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func generator() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

//Рекурсия --------------------------->
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func recursion() {
	x:= uint(5)
	res := factorial(x)
	fmt.Println(res)
}

//Panic ---------------------------->
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func panick() {
	defer second()
	first()
}

func panick2() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// tasks -------------------------->
