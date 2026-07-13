package main

import (
	"errors"
	"fmt"
)

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Division by 0 Error")
	}
	return num1 / num2, nil
}

func deferredFunction() {
	fmt.Println("Deferring...")
	r := recover()
	if r == nil {
		fmt.Println("No error occurred")
	} else {
		fmt.Println(r)
	}
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error occurred.")
	} else {
		fmt.Println("No error, Code executing further...")
		fmt.Println(result)
	}

	// defer deferredFunction()
	// panic("this caused a crash")

	// fmt.Println("Running...")

}
