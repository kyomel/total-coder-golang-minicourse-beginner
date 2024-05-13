package main

import "fmt"

// 1.2 BUILTIN TYPES
// structs
// map
// slice
// array

// type Player struct {
// 	name        string
// 	health      int
// 	attackPower int
// }

// func (player Player) Attack() {
// 	fmt.Printf("%s is attacking with attack power %d", player.name, player.attackPower)
// }

func main() {
	// structs
	// player := Player{
	// 	name:        "kyomel",
	// 	health:      100,
	// 	attackPower: 80,
	// }

	// another way to declare struct
	// player := Player{
	// 	"kyomel",
	// 	100,
	// 	80,
	// }
	// fmt.Println(player)

	// call by function
	// player.Attack()
	// fmt.Println(player.name)

	// map
	// balances := make(map[string]float64)
	// balances := map[string]float64{}
	// balances["michael"] = 100.0
	// fmt.Println(balances)
	// fmt.Println(balances["michael"])
	//if not exists
	// fmt.Println(balances["vania"])
	// another reach map
	// value, ok := balances["michael"]
	// if ok {
	// 	fmt.Println("the balance is ", value)
	// } else {
	// 	fmt.Println("doesn't exist at map")
	// }

	// slice
	// ints := []int{1, 2, 3, 4, 5}
	// ints = append(ints, 6)
	// fmt.Println(ints[2])
	// fmt.Println(ints[0:3])
	// otherInts := []int{6, 7, 8}
	// ints = append(ints, otherInts...)
	// another way slices
	// ints := make([]int, 10)
	// fmt.Println(ints)

	// array
	var ints [2]int
	ints[0] = 1
	ints[1] = 2
	fmt.Println(ints)
}
