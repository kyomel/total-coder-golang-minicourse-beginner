package main

import (
	"calculator/calculate"
	"fmt"
)

func main() {
	result := calculate.Add(2, 3)
	fmt.Println("the result is", result)
}
