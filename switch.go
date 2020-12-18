package main

import "fmt"

func main() {

	x := 1
	switch x {
	case 0:
		fmt.Print("0")
	case 1:
		fmt.Print("1")
	default:
		fmt.Print("default")
	}
}
