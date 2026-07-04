package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("--- Mathematical Constants ---")
	fmt.Printf("Pi (π): %f\n", math.Pi)
	fmt.Printf("Euler's Number (e): %f\n", math.E)
	fmt.Printf("Square root of 2: %f\n", math.Sqrt2)

	fmt.Println("\n--- Power & Root Functions ---")

	fmt.Printf("2^3 = %f\n", math.Pow(2, 3))

	fmt.Printf("Square root of 16 = %f\n", math.Sqrt(16))

	fmt.Printf("Cube root of 27 = %f\n", math.Cbrt(27))

	fmt.Println("\n--- Rounding Functions ---")
	val := 3.75

	fmt.Printf("Ceil(%f) = %f\n", val, math.Ceil(val))

	fmt.Printf("Floor(%f) = %f\n", val, math.Floor(val))

	fmt.Printf("Round(%f) = %f\n", val, math.Round(val))
	fmt.Printf("Round(3.25) = %f\n", math.Round(3.25))

	fmt.Printf("Trunc(%f) = %f\n", val, math.Trunc(val))

	fmt.Println("\n--- Extremes & Absolute Values ---")
	negNum := -15.45

	fmt.Printf("Absolute value of %f = %f\n", negNum, math.Abs(negNum))

	fmt.Printf("Max of 5.5 and 10.2 = %f\n", math.Max(5.5, 10.2))
	fmt.Printf("Min of 5.5 and 10.2 = %f\n", math.Min(5.5, 10.2))

	fmt.Println("\n--- Trig & Log Functions ---")

	rad90 := 90 * (math.Pi / 180)
	fmt.Printf("Sin of 90 degrees = %f\n", math.Sin(rad90))
	fmt.Printf("Cos of 90 degrees = %f\n", math.Cos(rad90))

	fmt.Printf("Natural Log of 10 = %f\n", math.Log(10))
	fmt.Printf("Log10 of 100 = %f\n", math.Log10(100))

	fmt.Println("\n--- Type Conversion (Important) ---")
	intBase := 5
	intExp := 2

	// This would fail: math.Pow(intBase, intExp)
	result := math.Pow(float64(intBase), float64(intExp))
	fmt.Printf("Int values converted to float64: 5^2 = %d (cast back to int)\n", int(result))
}
