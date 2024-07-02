package main

import (
	"calculator/calculate"
	"fmt"
)

func main() {
	player := calculate.Player{
		Name: "Vania",
	}
	result := calculate.Add(2, 3)
	fmt.Println("the result is", result)
	fmt.Println("the player name is", player.Name)
}
