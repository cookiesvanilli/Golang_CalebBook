package main

import (
	"fmt"
	"math"
)
func main() {
	shapeMain()
	nameMain()
}
type Shape interface {
	area() float64
}
type MultiShape struct {
	shapes []Shape
}
type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
//Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

//Rectangle
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//Shape

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
func shapeMain() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	m := MultiShape{[]Shape{&r, &c}}
	fmt.Println(m.area())
}

//-------------------->

type Person struct {
	Name string
	Address Address
}
type Address struct {
	City   string
}
type Android struct {
	Person
	NameMob string
	Model string
}
func (p *Person) Talk() {
	fmt.Println("My name is", p.Name)
}
func (p *Person) Location() {
	fmt.Println("From",  p.Address.City)
}
func (p *Android) Talk() {
	/*p.Person.Talk()*/
	fmt.Println("My phone model is", p.NameMob)
	fmt.Println("My model is", p.Model)
}
func nameMain() {
	p:= Person{Name: "Anna"}
	p.Talk()

	mobile := new(Android)
	mobile.NameMob = "Xiaomi"
	mobile.Model = "12345"
	mobile.Talk()

	p.Address.City = "Gotham"
	p.Location()
}
//tasks------------------>
//Какая разница между методом и функцией?
//Функция является независимой частью кода,
//связывающей один или несколько входных параметров с одним или
//несколькими выходными параметрами.
//Метод - это функция со специальным аргументом, известным как получатель.
//Получатель указывается в отдельном списке аргументов между
//ключевым словом func и именем метода.
/*Функция - это фрагмент кода, который вызывается по имени. Это может быть передано данных для работы (то есть параметров) и может при необходимости возвращать данные (возвращаемое значение).
Все данные, переданные функции, явно передаются.
A метод - это фрагмент кода, вызываемый именем, связанным с объектом.
В большинстве случаев он идентичен функции, за исключением двух ключевых отличий:
Метод неявно передает объект, на который он был вызван.
Метод способен работать с данными, содержащимися в классе
(помня, что объект является экземпляром класса - класс является определением, объект является экземпляром этих данных).
*/
