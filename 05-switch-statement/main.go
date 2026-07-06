package main

import "fmt"

func main() {

	// Basic Expression Switch
	finger := 3

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	default:
		fmt.Println("Unknown finger")
	}

	// Multiple Values in One Case
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Back to work.")
	}

	// Conditional Switch
	age := 21

	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}

	// Switch with Short Statement
	switch os := "linux"; os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other OS")
	}

	// Fallthrough Keyword
	num := 10

	switch num {
	case 10:
		fmt.Println("Matches 10")
		fallthrough
	case 20:
		// Executes because of fallthrough
		fmt.Println("Matches 20")
		fallthrough
	case 30:
		// Executes because of fallthrough
		fmt.Println("This prints too!")
	}

}
