package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	Name string
	Age int
}

func main() {
	persons := []person{}

	aom := person{Name : "aom", Age : 19};
	au := person{Name : "au", Age :20};

	persons = append(persons, aom);
	persons = append(persons, au);
	fmt.Println(persons)

	jsonData, err := json.Marshal(persons);
	
	if err != nil{
		fmt.Println(err);
	}
	
	fmt.Println(string(jsonData));


}