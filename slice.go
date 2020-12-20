package main

import "fmt"

func main() {
	name := []string{}
	var name2 []string

	name = append(name, "Test")
	name = append(name, "Test")
	name = append(name, "Test")
	name = append(name, "Test")

    name2 = append(name, "testest")
	fmt.Println(name)
	fmt.Println(name2)
}
