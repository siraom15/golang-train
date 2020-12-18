package main

import "fmt"

func main() {
	var arr [3]int64
	arr[0] = 1
	arr[1] = 3
	arr[2] = 4

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, v := range arr {
		fmt.Println(v)
	}

}