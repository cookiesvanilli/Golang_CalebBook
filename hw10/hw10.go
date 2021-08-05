package main

import (
	"fmt"
	"math/rand"
	"time"
)

//gorutine
func f(n int) {
	for i := 0; i < 12; i++ {
		fmt.Println(n, ":", i)
	}
}
func count() {
	go f(3)
	var input string
	fmt.Scanln(&input)
}
func start() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
func timing(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func startTime() {
	for i := 0; i < 10; i++ {
		go timing(i)
	}
	var input string
	fmt.Scanln(&input)
}
/*func main() {
	go spinner (25 * time.Millisecond)
	x := 45
	fmt.Printf("Fibonacci(#{x}) = #{fibonacci(x)}")

}
func spinner(duration time.Duration) {
	for {
		for _, v := range "-\\|/" {
			fmt.Printf("#{v}\r")
			time.Sleep(duration)
		}
	}
}
func fibonacci(x int) int {
	if x == 0 || x == 1 {
		fmt.Println(x)
		return  x
	}
	fmt.Println(fibonacci(x-1) + fibonacci(x-2))
	return fibonacci(x-1) + fibonacci(x-2)
}*/
//channel
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func channel() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
func main() {
	//start()
	//count()
	//startTime()
	channel()
}