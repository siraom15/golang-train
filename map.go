package main

import "fmt"

func main() {
	map1 := make(map[int]string)

	map1[0] = "hello"
	fmt.Println(map1)


	map2 := make(map[string]string)
	map2["person1"] = "aom"
	map2["person2"] = "mie"
	fmt.Println(map2)
	delete(map2, "person2")
	fmt.Println(map2)

}
