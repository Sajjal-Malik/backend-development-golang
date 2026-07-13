package main

import "fmt"

func main() {
	var a [5]int
	// Output: [0 0 0 0 0]
	fmt.Println(a)

	// Initialized with specific values
	b := [3]int{10, 20, 30}
	fmt.Println(b)

	// Size inferred by the compiler using ellipsis (...)
	c := [...]string{"Go", "Zig", "Rust"}
	fmt.Println(c)

	// Initializing specific indices (index 1 and 3, others get 0)
	d := [5]int{1: 12, 3: 45}
	fmt.Println(d)

	primes := [4]int{2, 3, 5, 7}

	// Getting the total size
	fmt.Println(len(primes))

	// Modifying an element
	primes[2] = 99

	// Iterating with both Index and Value
	for index, value := range primes {
		fmt.Printf("Index %d has value %d\n", index, value)
	}

	// Iterating and ignoring the index using a blank identifier (_)
	for _, value := range primes {
		fmt.Println(value)
	}

	// Declares a 2x3 integer matrix (2 rows, 3 columns)
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	// Accesses row 1, column 2. Output: 6
	fmt.Println(matrix[1][2])

}
