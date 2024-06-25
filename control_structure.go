package main

import (
	"fmt"
)

// 1.6 CONTROL STRUCTURES
// - for loops
// - for range
// - switch

func main() {
	// names := []string{"francesca", "vania", "asteria", "ayu"}
	// for loops
	// for i := 0; i < len(names); i++ {
	// 	// fmt.Println("from inside the loop", i)
	// 	fmt.Println("the names is", names[i])
	// }
	// for range
	// for index, name := range names {
	// 	fmt.Printf("this is the index of the slice %d and the name is %s\n", index, name)
	// }

	// balances := map[string]int{
	// 	"francesca": 100,
	// 	"vania":     2,
	// 	"asteria":   1000,
	// 	"ayu":       2,
	// }

	// for key, value := range balances {
	// 	fmt.Printf("%s has balance of %d\n", key, value)
	// }

	// switch
	status := "accepted"

	switch status {
	case "accepted":
		fmt.Println("status is accepted")
	case "pending":
		fmt.Println("status is pending")
	default:
		fmt.Println("status is unknown")
	}
}
