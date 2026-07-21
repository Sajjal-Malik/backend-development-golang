package main

import "fmt"

func change(i *int) {
	// fmt.Println("Before change i is: ", *i)
	*i = 1000
	// fmt.Println("After change i is: ", *i)
}

type Book struct {
	id    int
	title string
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func test(pointerSlice *[]*int) []*int {
	values := *pointerSlice

	for range values {
	}

	return values
}

func copyToNewSlice(pointerSlice *[]*int) []*int {
	originalValues := *pointerSlice

	// Initialize an empty slice
	var newSlice []*int

	// Loop and append the values
	for _, value := range originalValues {
		newSlice = append(newSlice, value)
	}

	// Return the new slice
	return newSlice
}

func copyToNewMap(pointerSlice *[]*int) map[int]*int {
	originalValues := *pointerSlice

	// Initialize an empty map using make()
	newMap := make(map[int]*int)

	// Loop and assign index as key, pointer as value
	for index, value := range originalValues {
		newMap[index] = value
	}

	// Return the new map
	return newMap
}

func main() {
	// var x int = 12
	// var y int = &x

	x := 12
	y := &x
	*y = 100
	// fmt.Println(x, *y)

	num := 10
	change(&num)
	// fmt.Println("Param value after func call: ", num)

	b := Book{10, "Old"}
	b.setTitle("New")
	// fmt.Println(b)

	num1 := 4
	num2 := &num1
	num3 := &num2

	fmt.Printf("Type of num1: %T\n", num1)
	fmt.Printf("Type of num2: %T\n", num2)
	fmt.Printf("Type of num3: %T\n", num3)

	fmt.Println("Num 1 value: ", num1)
	fmt.Println("Num 2 value: ", *num2)
	fmt.Println("Num 3 value: ", **num3)

	num4 := 10
	num5 := 20

	values := &[]*int{&num4, &num5}
	fmt.Println(test(values))

	// Receives and prints the new slice
	resultSlice := copyToNewSlice(values)
	// Prints hexadecimal pointer addresses
	fmt.Println(resultSlice)

	// Receives and prints the new map
	resultMap := copyToNewMap(values)
	// Prints hexadecimal pointer addresses
	fmt.Println(resultMap)

}
