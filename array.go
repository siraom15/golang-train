package main

import "fmt"

func main() {
	var arr [3]int64
	arr[0] = 1
	arr[1] = 3
	arr[2] = 4

	arr2 := [5]int64{1,2,3,5,6} //declare w/ value

	//for
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//for-range
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Print(arr)
}