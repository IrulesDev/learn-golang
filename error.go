package main

import (
	"errors"
	"fmt"
)

func main(){
	res, err :=findItem(1)
	if error.is (err, errorNotFound){
		fmt.println("eitem not found")

	}
	var errorNotFound = errors.New("not found")
	
	fmt.Println(res)
	func findItem(id int)(string error){
		if id = 1{
			return "", errors.New("error")
		}
		return "item Found",nil
	}
}

