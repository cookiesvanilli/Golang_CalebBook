package main

import "fmt"

func main() {
	five()
	pointerFive()
	oneNew()
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
