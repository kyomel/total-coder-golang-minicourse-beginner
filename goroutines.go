package main

func DoWork() {
	result := "the result of do work"
	_ = result
}

func main() {
	// Run Style 1
	go DoWork()
	// Run Style 2
	// go func() {
	// 	fmt.Println("we are doing some work!")
	// }()
	// time.Sleep(time.Second)
	// exit
}
