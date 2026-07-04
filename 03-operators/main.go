package main

import "fmt"

func main() {

	fmt.Println("--- Arithmetic Operators ---")
	a, b := 10, 3

	fmt.Printf("%d + %d = %d (Addition)\n", a, b, a+b)
	fmt.Printf("%d - %d = %d (Subtraction)\n", a, b, a-b)
	fmt.Printf("%d * %d = %d (Multiplication)\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (Division/Integer Truncation)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (Modulus/Remainder)\n", a, b, a%b)

	c := 5
	c++
	fmt.Printf("c++ = %d (Increment)\n", c)
	c--
	fmt.Printf("c-- = %d (Decrement)\n", c)

	fmt.Println("\n--- Relational Operators ---")
	x, y := 12, 20

	fmt.Printf("%d == %d is %t (Equal to)\n", x, y, x == y)
	fmt.Printf("%d != %d is %t (Not equal to)\n", x, y, x != y)
	fmt.Printf("%d > %d is %t (Greater than)\n", x, y, x > y)
	fmt.Printf("%d < %d is %t (Less than)\n", x, y, x < y)
	fmt.Printf("%d >= %d is %t (Greater than or equal to)\n", x, y, x >= y)
	fmt.Printf("%d <= %d is %t (Less than or equal to)\n", x, y, x <= y)

	fmt.Println("\n--- Logical Operators ---")
	t, f := true, false

	fmt.Printf("%t && %t is %t (Logical AND)\n", t, f, t && f)

	fmt.Printf("%t || %t is %t (Logical OR)\n", t, f, t || f)

	fmt.Printf("!%t is %t (Logical NOT)\n", t, !t)

	fmt.Println("\n--- Bitwise Operators ---")
	// p = 60 (binary: 0011 1100), q = 13 (binary: 0000 1101)
	p, q := 60, 13

	fmt.Printf("%d & %d = %d (Bitwise AND: sets bit to 1 if both bits are 1)\n", p, q, p&q)
	fmt.Printf("%d | %d = %d (Bitwise OR: sets bit to 1 if one of the bits is 1)\n", p, q, p|q)
	fmt.Printf("%d ^ %d = %d (Bitwise XOR: sets bit to 1 if only one bit is 1)\n", p, q, p^q)
	fmt.Printf("%d &^ %d = %d (Bitwise AND NOT / Bit Clear: clears bits of p if q's bit is 1)\n", p, q, p&^q)
	fmt.Printf("%d << 2 = %d (Left Shift: shifts bits left by 2 positions)\n", p, p<<2)
	fmt.Printf("%d >> 2 = %d (Right Shift: shifts bits right by 2 positions)\n", p, p>>2)

	fmt.Println("\n--- Assignment Operators ---")
	var z int

	z = 10 // Simple assignment
	fmt.Printf("z = %d\n", z)

	z += 5 // Same as: z = z + 5
	fmt.Printf("z += 5 gives %d\n", z)

	z -= 3 // Same as: z = z - 3
	fmt.Printf("z -= 3 gives %d\n", z)

	z *= 2 // Same as: z = z * 2
	fmt.Printf("z *= 2 gives %d\n", z)

	z /= 4 // Same as: z = z / 4
	fmt.Printf("z /= 4 gives %d\n", z)

	z %= 2 // Same as: z = z % 2
	fmt.Printf("z %%= 2 gives %d\n", z)

	z = 5
	z <<= 1 // Same as: z = z << 1
	fmt.Printf("z <<= 1 gives %d\n", z)

	fmt.Println("\n--- Miscellaneous Operators ---")

	num := 42
	ptr := &num
	fmt.Printf("Address of num (&num): %v\n", ptr)

	fmt.Printf("Value at ptr (*ptr): %d\n", *ptr)

	// Channel Operator (<-)
	ch := make(chan string, 1)
	ch <- "Hello Go Channel" // Send value into channel
	msg := <-ch              // Receive value from channel
	fmt.Printf("Channel data (<-ch): %s\n", msg)
}
