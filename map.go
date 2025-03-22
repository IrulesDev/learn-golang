package main

import "fmt"

func main (){

	// var person map[string] string{}
	// person["nama"] = "irul"
	// person["adress"] = "gantiwarno"

	person := map[string] string{
		"name" : "irul",
		"address" : "gantiwarno",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
}