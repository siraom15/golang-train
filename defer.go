package main

import "fmt"

func main() {
	defer end()
	defer end2()
	fmt.Println("Hello World")
}

func end() {
	fmt.Println("End")
}
func end2(){
	fmt.Println("End2")
}