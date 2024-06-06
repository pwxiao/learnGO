package main

import "fmt"

const (
	BEIJING = iota
	SHANGHAI
	NANJING
)

func main() {

	const length = 10
	fmt.Println("BEJING IS ", BEIJING)
	fmt.Println("SHANGHAI IS ", SHANGHAI)
	fmt.Println("NANJING IS ", NANJING)
	fmt.Println("Length is ", length)
}
