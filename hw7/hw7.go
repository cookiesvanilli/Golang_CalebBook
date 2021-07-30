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
	task1()
	fmt.Println(" ---------------")
	task2()
	max(12, 10)
	task4()
	fmt.Println("fibonacci ---------------")
	fib()
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
func task2() {
	number:= 35
	half := number / 2
	if (half % 2) == 0 {
		fmt.Println(half, true)
	} else {
		fmt.Println(half, false)
	}
}
func max(x, y int64) int64 {
	if (x < y) {
		fmt.Println(y)
		return y
	}
	fmt.Println(x)
	return x
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
func fibonacci() func() int {
	numbers := make(map[int]int)
	n := 0
	return func() int {
		if n == 0 {
			numbers[n] = 0
			n++
			return 0
		}
		if n == 1 {
			numbers[n] = 1
			n++
			return 1
		}
		number := numbers[n-1] + numbers[n-2]
		numbers[n] = number
		n++
		return number
	}
}
func fib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
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