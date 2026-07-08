package main

import "fmt"

// Basic function taking two integers and returning an integer
func add(x int, y int) int {
	return x + y
}

// Using Shortened Parameter Types
func multiply(x, y, z int) int {
	return x * y * z
}

//  Multiple Return Values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Defining function as parameters
func callFunc(callable func(int) int) int {
	return callable(10)
}

func doubleNumber(number int) int {
	return number * 2
}

func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

// Variadic Functions
func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Named Return Values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// automatically returns x and y
	return
}

func main() {
	// result := add(10, 5)
	// fmt.Println("Sum:", result)

	// value := callFunc(doubleNumber)

	f1 := func(number int) int {
		return number * 2
	}

	value := callFunc(f1)
	fmt.Println(value)
	fmt.Printf("%T \n", value)
	fmt.Printf("%T \n", doubleNumber)

	f := getFunc("hello ")
	res := f("world")
	fmt.Println(res)
	fmt.Printf("%T \n", res)
	fmt.Printf("%T \n", f)

	// We are passing slice here
	result1 := sumAll(3, 4, 5, 6)

	// More explicit way of passing slice
	result2 := sumAll([]int{3, 4, 5, 6}...)

	fmt.Println(result1)
	fmt.Println(result2)

	var first, second int
	first, second = split(40)
	fmt.Println(first, second)
}
