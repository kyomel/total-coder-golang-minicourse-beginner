package main

import (
	"calculator/calculate"
	"fmt"
)

func main() {
	result := calculate.Add(10, 11)
	fmt.Println("hello from the printname app", result)
}
