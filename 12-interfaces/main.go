package main

import "fmt"

type Shape interface {
	getPerimeter() uint
}

type Square struct {
	width uint
}

func (S Square) getPerimeter() uint {
	return 4 * S.width
}

type Triangle struct {
	a uint
	b uint
	c uint
}

func (T Triangle) getPerimeter() uint {
	return T.a + T.b + T.c
}

func (T Triangle) getSides() []uint {
	return []uint{T.a, T.b, T.c}
}

func isEventPerimeter(shape Shape) bool {
	return shape.getPerimeter()%2 == 0
}

func main() {
	// var s Shape = Triangle{1, 2, 3}
	// fmt.Println(s.getPerimeter())

	perimeter := uint(0) // type-casting

	var shapes []Shape = []Shape{Triangle{1, 2, 3}, Square{10}}
	for _, shape := range shapes {
		perimeter += shape.getPerimeter()
	}
	fmt.Println(perimeter)
}
