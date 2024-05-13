package main

import "fmt"

// 1.1 TYPE SYSTEM

// variables in the global space
// variables in the local space

// global space
var (
	age                   = 30        //int
	unassignedAge uint    = 30        //uint
	name                  = "michael" // string
	balance               = 4.44      // float64
	balance32     float32 = 4.44      // float32
)

const (
	foo     = "foo"
	timeout = 10
)

func main() {
	// local space
	// Set and Declare Golang
	// email := "michael@mail.com"
	// attackPower := 100
	// health := 99.0

	// var health float64 // 0.0
	// var str string     // ""
	// str = "michael"
	// fmt.Println(health)
	// fmt.Println(str)

	attackPower := 100 // int
	health := 200.0    // float64
	fmt.Println(health - float64(attackPower))
}
