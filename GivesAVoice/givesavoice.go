package main

import (
	"fmt"
)

type Givesavoice interface {
	SayHelloLoud() string
	SayHelloQuiet() string
}
type Louder struct {
	voices []Givesavoice
}
type Quieter struct {
	voices []Givesavoice
}
type Cat struct {
	Greeting string
}
type Dog struct {
	Greeting string
}
type Bird struct {
	Greeting string
}
type Bee struct {
	Greeting string
}
type Cow struct {
	Greeting string
}
func (c *Cat) SayHelloLoud() string {//Cat
	fmt.Println("Cat: MEOW")
	return "Cat: MEOW"
}
func (cq *Cat) SayHelloQuiet() string {
	fmt.Println("Cat: meow")
	return "Cat: meow"
}
func (d *Dog) SayHelloLoud() string {//Dog
	fmt.Println("Dog: WOFF")
	return "Dog: WOFF"
}
func (b *Bird) SayHelloLoud() string {//Bird
	fmt.Println("Bird: WEE-TWEET-TWEET")
	return "Bird: WEE-TWEET-TWEET"
}
func (e *Bee) SayHelloLoud() string {//Bird
	fmt.Println( "Bee: BZZZ")
	return "Bee: BZZZ"
}
//Cow
func (o *Cow) SayHelloLoud() string {
	fmt.Println( "Cow: MOO")
	return "Cow: MOO"
}
//Volume
func (l *Louder) SayHelloLoud() string {
	var shout string
	for _, s := range l.voices {
		shout += s.SayHelloLoud()
	}
	return shout
}
func (q *Quieter) SayHelloQuite() string {
	var shout string
	for _, s := range q.voices {
		shout += s.SayHelloQuiet()
	}
	return shout
}
func main() {
c:= Cat{"Cat: MEOW"}
c.SayHelloLoud()
d:= Dog{ "Dog: WOFF"}
d.SayHelloLoud()
b:= Bird{ "Bird: WEE-TWEET-TWEET"}
b.SayHelloLoud()
e:= Bee{"Bee: BZZZ"}
e.SayHelloLoud()
o:= Cow{"MOO"}
o.SayHelloLoud()
cq:=Cat{"meow"}
cq.SayHelloQuiet()
}