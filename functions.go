package main

import "fmt"

// 1.3 FUNCTIONS

func myFunctions(name string, age int) (int, string) {
	return age, name
}

func main() {
	age, name := myFunctions("Vania", 25)
	fmt.Println(age, name)
}
