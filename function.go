package main

import "fmt"

func main() {
	var number1 float64
	var number2 float64
	fmt.Print("Input Number 1 :")
	fmt.Scanf("%f",&number1)
	fmt.Print("Input Number 2 :")
	fmt.Scanf("%f",&number2)

	fmt.Println("Result is",addition(number1, number2))
}

func addition(a float64, b float64) float64 {
	return a+b
}
