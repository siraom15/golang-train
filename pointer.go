package main

import "fmt"

func main() {

	var point *int
	var myNum = 630
	point = &myNum
	fmt.Println(*point)
	*point = 300
	fmt.Println(*point)
	fmt.Println(myNum)
}
