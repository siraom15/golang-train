package main

func main() {
	// declare value
	var myInt int64 // default val = 0
	var myStr string // default val = ""
	var myFloat float64 // default val = 0.00
	var myBool bool // default val = false

	// declare w/ value

	var myTestInt int64 = 1
	var myTestStr string = "Hello World"
	var myTestFloat float64 = 43.23
	var myTestBool bool = true

	// short declare auto detect type
	shortInt := 1
	shortStr := "Hello World"
	shortFloat := 3123.3
	shortBool := true


	//
	var(
		a = 3
		b = 4
		c = "hello"
	)
	
	//constant 
	const secret string = "IloveCat"

	// secret = "hello" // will error 
}