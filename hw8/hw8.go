package main

import "fmt"

func main() {
	five()
	pointerFive()
	oneNew()
	fmt.Println("------>")
	task1()
	change()
}
func zero(x int) {
	x = 0
}
func five() {
	x := 5
	zero(x)
	fmt.Println(x) // x всё еще равен 5
}
func pointerZero(xPtr *int) {
	*xPtr = 0
}
func pointerFive() {
	x := 5
	pointerZero(&x)
	fmt.Println(x) // x is 0
}
//Operator new --------------->
func one(xPtr *int) {
	*xPtr = 1
}
func oneNew() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
//tasks ---------------->
//Как получить адрес переменной?
//Существует оператор &, который используется для получения адреса переменной.
//--------->
//Как присвоить значение указателю?
//В Go указатели представлены через оператор * (звёздочка),
//за которым следует тип хранимого значения.
//В функции zero xPtr является указателем на int.
//* также используется для «разыменовывания» указателей.
//Также существует оператор &, который используется для получения адреса переменной.
//--------->
//Как создать новый указатель?
//Функция new принимает аргументом тип, выделяет для него память и возвращает указатель на эту память.
//-------->
func square(x *float64) {
	*x = *x * *x
}
func task1() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}
//-------->
func swap(x, y *int) (int, int){
	*x, *y = *y, *x
	return *x, *y
}
func change() {
	x, y := 1, 2
	swap(&x,&y)
	fmt.Println("x =", x,"y =", y)
}