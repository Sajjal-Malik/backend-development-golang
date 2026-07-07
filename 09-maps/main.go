package main

import "fmt"

func main() {

	// Nil map (cannot write to this)
	var nilMap map[string]int

	// Initializing with make()
	ages := make(map[string]int)
	ages["Alice"] = 30

	// Initializing with a Map Literal
	scores := map[string]int{
		"Go":   100,
		"Java": 85,
		"C#":   85,
	}

	// Add or Update
	scores["Go"] = 98

	// Read with existence check (Comma OK)
	val, exists := scores["Python"]
	if exists {
		fmt.Println("Python score:", val)
	} else {
		fmt.Println("Python key not found.")
	}

	// Delete an element form map
	delete(scores, "Java")

	// Check total items
	fmt.Println("Total items:", len(scores))

	fmt.Println(ages, scores, nilMap == nil)

	mp := map[uint]uint{}
	n := uint(100)

	for number := uint(1); number <= n; number++ {
		for d := uint(1); d <= 5; d++ {
			if number%d == 0 {
				mp[d]++
			}
		}
	}
}
