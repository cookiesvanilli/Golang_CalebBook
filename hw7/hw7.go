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
	deferredCall()
	panick2()
	task1()
	fmt.Println(" ---------------")
	task2()
	task3()
	task4()
	fmt.Println("fibonacci ---------------")
	fibonacci(10)
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
// Closures ------------------------->
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
//Recursion --------------------------->
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
func deferredCall() {
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
func task1() {
	numb := []int{1, 2, 3, 4, 5}
	fmt.Print("Result: ", sum(numb))
}
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
//half------------------->
func task2() {
	fmt.Println(half(20))
	fmt.Println(half(21))
}
func half(n int) (int, bool) {
	return n / 2, n % 2 == 0
}
//max ------------------->
func task3() {
	fmt.Println (max(33,55,66,77,852,654,158))
}
func max (x ...int)int {
	var y int
	for _ , d := range x {
		if d > y {
			y = d
		}
	}
	return y
}
func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i + 1
		i += 2
		return
	}
}
func task4() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
func fibonacci(x int) int {
if x == 0 || x == 1 {
	fmt.Println(x)
	return  x
}
	fmt.Println(fibonacci(x-1) + fibonacci(x-2))
return fibonacci(x-1) + fibonacci(x-2)

}

//Что такое отложенный вызов, паника и восстановление?
//Как восстановить функцию после паники?
//В Go есть специальный оператор defer,
//который позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая.
//Ранее мы создали функцию, которая вызывает panic, чтобы сгенерировать ошибку выполнения.
//Мы можем обрабатывать паники с помощью встроенной функции recover.
//Функция recover останавливает панику и возвращает значение,
//которое было передано функции panic.
//Паника обычно указывает на ошибку программиста (например, попытку получить доступ
//к несуществующему индексу массива, забытая и непроинициализированная карта и т.д.)
//или неожиданное поведение (исключение), которое нельзя обработать
//(поэтому оно и называется «паника»).