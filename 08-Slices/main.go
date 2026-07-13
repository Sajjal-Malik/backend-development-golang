package main

import "fmt"

func main() {
	// Slice Literal
	s1 := []int{10, 20, 30}
	fmt.Println(s1)

	// Using the make() function => Syntax: make([]Type, length, capacity)
	s2 := make([]int, 3, 5)
	fmt.Println(s2)

	// Slicing an existing array or slice
	arr := [5]int{1, 2, 3, 4, 5}
	// Grabs elements from index 1 to 3 -> [2, 3, 4]
	s3 := arr[1:4]
	fmt.Println(s3)

	var s []int
	// s is now [1]
	s = append(s, 1)
	// s is now [1, 2, 3] (can append multiple items)
	s = append(s, 2, 3)

	for index, val := range s {
		fmt.Printf("Index: %d, Value: %d\n", index, val)
	}

	fruitsArray := [4]string{"Apple", "Banana", "Cherry", "Date"}

	fruitsSlice := fruitsArray[1:3]
	fmt.Println(fruitsSlice)

	colors := []string{"Red", "Green", "Blue"}

	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

}
