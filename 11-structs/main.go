package main

import "fmt"

// Structs
type Sport struct {
	name     string
	position string
}

type Person struct {
	name string
	age  int
	// showName  func(string) string
	favtSports []Sport
}

func getName(p Person) string {
	return p.name
}

func getPerson(p Person) Person {
	return p
}

// Methods on struct
func (p Person) getAge() int {
	return p.age
}

func (p Person) setAge(age int) {
	p.age = age
	fmt.Println(p)
}

func main() {

	// p1 := Person{"Malik", 26}
	// p1 := Person{name: "Malik", age: 26}
	// fmt.Println(p1)

	// var p1 Person = Person{age: 26}
	p1 := Person{age: 26}
	p1.name = "Malik"
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	fmt.Println(getName(p1))
	fmt.Println(getPerson(p1))

	// p1.showName = func(name string) string {
	// 	return name
	// }
	// fmt.Println(p1.showName("Bhatti"))

	value := p1.getAge()
	fmt.Println(value)

	p1.setAge(28)
	fmt.Println(value)

	p2 := Person{name: "Malik", age: 26, favtSports: []Sport{{"Football", "Attacker"}}}
	fmt.Println(p2)
	fmt.Println(p2.favtSports[0])
	fmt.Println(p2.favtSports[0].name)
	fmt.Println(p2.favtSports[0].position)

}
