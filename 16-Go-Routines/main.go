package main

import (
	"fmt"
	"time"
)

func run1() {
	time.Sleep(2 * time.Second)
	fmt.Println("print run 1!")
}

func run2() {
	time.Sleep(4 * time.Second)
	fmt.Println("print run 2!")
}

func run3() {
	time.Sleep(6 * time.Second)
	fmt.Println("print run 3!")
}

func add(num1 int, num2 int, channel chan int) {
	// time.Sleep(5 * time.Second)
	channel <- num1 + num2
	// return num1 + num2
}

func newAdd(num1 int, num2, delay int, channel chan int) {
	time.Sleep(time.Duration(delay) * time.Second)
	channel <- num1 + num2
}

func main() {
	go run1()
	go run2()
	go run3()
	// time.Sleep(7 * time.Second)
	fmt.Println("All Routines Done!")

	// This will panic, Can't store returned Values from GO Routine
	// ans := go add(4, 6)
	// fmt.Println("Add Done!")

	ch := make(chan int)
	go add(4, 5, ch)
	ans := <-ch
	go add(12, 23, ch)
	ans = <-ch
	go add(33, 43, ch)
	ans = <-ch
	go add(99, 199, ch)
	ans = <-ch
	go add(45, 155, ch)
	ans = <-ch
	fmt.Println(ans)
	fmt.Println("Go Routine with Channel, Done!")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go newAdd(20, 30, 5, ch1)
	go newAdd(40, 50, 7, ch2)

	res1 := <-ch1
	res2 := <-ch2

	fmt.Println(res1)
	fmt.Println(res2)
}
