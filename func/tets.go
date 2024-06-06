package main

import "fmt"

func foo1(a string, b int) (string, int) {
	return a, b
}

func foo2(a string, b int) (ret1 string, ret2 int) {
	ret1 = a
	ret2 = b
	return
}
func main() {
	ret1, ret2 := foo2("nmsl", 111)
	fmt.Println("ret1 is ", ret1, "ret2 is ", ret2)
}
