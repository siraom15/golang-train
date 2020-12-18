package main

import "fmt"

func main() {
	var aom student = student{"Siriwat", 24, 4.00, true}
	fmt.Println(aom)

	person := struct {
		name    string
		age     int16
		career  string
		isAlive bool
	}{
		name:    "none",
		age:     0,
		career:  "none",
		isAlive: true,
	}
	var aom2 = person
	fmt.Println(aom2)
}

//zero value
type student struct {
	name      string
	age       int16
	gpa       float32
	isStudent bool
}

//set default value

