package main

import "fmt"

// 1.4 CUSTOM TYPES

// a
// type mySpecialAge int
// func printAge(age int) {
// 	fmt.Println(age)
// }
// func printSpecialAge(age mySpecialAge) {
// 	fmt.Println(age)
// }

// b
// type specialAge int
// func (age specialAge) dooStuff() {
// 	fmt.Println("this is my special age:", age)
// }

// c
type specialFunc func(int) int

func foo(fn specialFunc) {
	result := fn(100)
	fmt.Println("result of my special func", result)
}

func myFunction(age int) int {
	return age
}

func main() {
	// a
	// var age int = 50
	// var specialAge mySpecialAge = 50
	// printAge(age)
	// printSpecialAge(specialAge)
	// another way
	// printSpecialAge(mySpecialAge(age))

	// b
	// var age specialAge = 50
	// age.dooStuff()

	// c
	foo(myFunction)
}
