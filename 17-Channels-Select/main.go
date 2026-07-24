package main

import (
	"fmt"
	"time"
)

// func add(num1 int, num2, delay int, channel chan int) {
// 	time.Sleep(time.Duration(delay) * time.Second)
// 	// Will NOT block because the channel has a buffer slot!
// 	channel <- num1 + num2
// 	fmt.Println("Goroutine finished writing and exited.")
// }

// channel chan<- int means this function can ONLY send data
func add(num1 int, num2, delay int, channel chan<- int) {
	time.Sleep(time.Duration(delay) * time.Second)
	channel <- num1 + num2
}

// receiveOnly can ONLY read data from the channel
func receiveOnly(channel <-chan int) {
	ans := <-channel
	fmt.Println("Received:", ans)

}

type Counter struct {
	value int
}

func count(counter *Counter) {
	counter.value++
	fmt.Println(counter.value)
}

func main() {
	ch := make(chan int, 1)
	ch2 := make(chan int, 1)

	go add(10, 20, 2, ch)
	go add(10, 20, 2, ch2)

	// ans := <-ch
	// ans2 := <-ch2
	// fmt.Println(ans, ans2)

	// for i := 1; i <= 2; i++ {
	// 	select {
	// 	case ans := <-ch:
	// 		fmt.Println(ans)
	// 	case ans2 := <-ch2:
	// 		fmt.Println(ans2)
	// 	}
	// }

	// Main sleeps longer than the goroutines
	time.Sleep(3 * time.Second)

	select {
	case ans := <-ch:
		fmt.Println("Main read:", ans)
	case ans2 := <-ch2:
		fmt.Println("Main read:", ans2)
	}

	// Main creates a bidirectional channel
	channel := make(chan int)

	// Passed as write-only to add
	go add(10, 20, 1, channel)
	receiveOnly(ch)

	// Create a buffered channel that can hold up to 3 booleans
	statusReports := make(chan bool, 3)

	// Send 3 boolean values into the channel
	// This does NOT block because the buffer size is 3
	statusReports <- true
	statusReports <- false
	statusReports <- true

	// Read and print the boolean values back out (First In, First Out)
	fmt.Println(<-statusReports)
	fmt.Println(<-statusReports)
	fmt.Println(<-statusReports)

	counter := Counter{0}

	for i := 1; i <= 100; i++ {
		go count(&counter)
	}
	time.Sleep(2 * time.Second)
}
