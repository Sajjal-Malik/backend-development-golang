package main

import "fmt"

// Without Generic
func add(num1 int, num2 int) int {
	return num1 + num2
}

type Number interface{ int | float64 | uint }

// T is a 'generic' type parameter.
func newAdd[T Number](num1 T, num2 T) T {
	var sum T = num1 + num2
	return sum
}

func getValues[K comparable, V any](mp map[K]V) []V {
	values := []V{}

	for _, value := range mp {
		values = append(values, value)
	}

	return values
}

// Declare the generic slice
type GenericSlice[T any] []T

type GenericStruct[T any, K any] struct {
	Values T
}

func main() {

	value1 := newAdd(4, 5)
	fmt.Println(value1)
	value2 := newAdd(4.55, 8.99)
	fmt.Println(value2)

	// var mp map[string]int = map[string][int]{"one":100, "two":200, "three": 300}
	mp := map[string]int{"one": 100, "two": 200, "three": 300}
	values := getValues(mp)
	fmt.Println(values)

	intSlice := GenericSlice[int]{1, 2, 3}
	fmt.Println(intSlice)

	strSlice := GenericSlice[string]{"apple", "banana"}
	fmt.Println(strSlice)

	instance1 := GenericStruct[int, string]{Values: 42}
	fmt.Println(instance1.Values)

	instance2 := GenericStruct[string, bool]{Values: "Hello, Go!"}
	fmt.Println(instance2.Values)
}
