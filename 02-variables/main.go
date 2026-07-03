package main

import "fmt"

func main() {
	// ==========================================
	// STRING TYPE (Text data)
	// ==========================================
	var courseName string = "Golang for Beginners"
	instructor := "Alice" // Short declaration (infers string)

	// ==========================================
	// INTEGER TYPES (Whole numbers)
	// ==========================================
	var positiveAndNegative int = -15
	// var positiveOnly uint = 50
	studentsCount := 120 // Short declaration (infers int)

	// ==========================================
	// FLOATING-POINT TYPES (Decimal numbers)
	// ==========================================
	var pi float32 = 3.14
	rating := 4.8 // Short declaration (infers float64)

	// ==========================================
	// BOOLEAN TYPE (True / False)
	// ==========================================
	var isLessonComplete bool = false
	isPassed := true // Short declaration (infers bool)

	// ==========================================
	// DEFAULT "ZERO VALUES"
	// (Variables declared without an explicit value)
	// ==========================================
	var defaultInt int       // Defaults to 0
	var defaultFloat float64 // Defaults to 0.0
	var defaultBool bool     // Defaults to false
	var defaultString string // Defaults to "" (empty string)

	// ==========================================
	// PRINTING THE VALUES
	// ==========================================
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
