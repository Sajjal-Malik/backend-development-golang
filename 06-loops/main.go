package main

import "fmt"

func main() {

	// Three-Component Loop (Traditional)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Condition-Only Loop (While Loop)
	count := 1
	for count <= 3 {
		fmt.Println(count)
		count++
	}

	// Infinite Loop
	// for {
	// 	fmt.Println("Looping...")
	// 	break
	// }

	// Range Loop (Foreach)
	numbers := []int{10, 20, 30}

	// Iterating over a slice
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Using blank identifier (_) if you only need the value
	for _, value := range numbers {
		fmt.Println(value)
	}

	// Labeled Break and Continue
	// OuterLoop:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 0; j < 3; j++ {
	// 			if i == 1 && j == 1 {
	// 				// Exits both loops instantly
	// 				break OuterLoop
	// 			}
	// 			fmt.Printf("i: %d, j: %d\n", i, j)
	// 		}
	// 	}

	// Looping Over String (Strings are stored as Arrays BTS)
	str := "Hello World"

	// for idx := 0; idx <= len(str); idx++ {
	// 	fmt.Printf("%c", str[idx])
	// }

	for _, char := range str {
		fmt.Printf("%c", char)
	}
}
