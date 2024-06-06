package main

import "fmt"

// 1.5 ENUMS

type Status int

func (status Status) String() string {
	switch status {
	case StatusCreated:
		return "CREATED"
	case StatusCanceled:
		return "CANCELED"
	case StatusPending:
		return "PENDING"
	case StatusFailed:
		return "FAILED"
	default:
		return "UNKNOWN"
	}
}

const (
	StatusCreated Status = iota // 0
	// below
	StatusCanceled
	StatusPending
	StatusFailed
)

func main() {
	handleStatus(StatusCanceled)
	// fmt.Println(StatusCreated)
	// fmt.Println(StatusCanceled)
	// fmt.Println(StatusPending)
	// fmt.Println(StatusFailed)
}

func handleStatus(status Status) {
	fmt.Println("we are handling the status: ", status)
	// fmt.Printf("we are handling the status: %s", status)
}
