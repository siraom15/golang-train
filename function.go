package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var number1 float64
	// var number2 float64
	// fmt.Print("Input Number 1 :")
	// fmt.Scanf("%f", &number1)
	// fmt.Print("Input Number 2 :")
	// fmt.Scanf("%f", &number2)

	// fmt.Println("Result is", addition(number1, number2))

	res, err := divide(0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(res)
}

func addition(a float64, b float64) (float64, error) {
	return a + b, nil
}
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("Cannot Devide by zero")
		return 0.0, err

	}
	return a / b, nil
}
