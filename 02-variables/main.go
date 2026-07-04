package main

import "fmt"

func main() {
	var courseName string = "Golang for Beginners"
	instructor := "Alice"

	var positiveAndNegative int = -15
	// var positiveOnly uint = 50
	studentsCount := 120

	var pi float32 = 3.14
	rating := 4.8

	var isLessonComplete bool = false
	isPassed := true

	var defaultInt int       // Defaults to 0
	var defaultFloat float64 // Defaults to 0.0
	var defaultBool bool     // Defaults to false
	var defaultString string // Defaults to "" (empty string)

	fmt.Println("--- String Examples ---")
	fmt.Printf("Course: %s, Instructor: %s\n\n", courseName, instructor)

	fmt.Println("--- Integer Examples ---")
	fmt.Printf("Score: %d, Total Students: %d\n\n", positiveAndNegative, studentsCount)

	fmt.Println("--- Float Examples ---")
	fmt.Printf("Pi: %.2f, Rating: %.1f\n\n", pi, rating)

	fmt.Println("--- Boolean Examples ---")
	fmt.Printf("Completed: %t, Passed: %t\n\n", isLessonComplete, isPassed)

	fmt.Println("--- Default Zero Values ---")
	fmt.Printf("Int: %d\n", defaultInt)
	fmt.Printf("Float: %.1f\n", defaultFloat)
	fmt.Printf("Bool: %t\n", defaultBool)
	fmt.Printf("String: %q\n", defaultString)
}
