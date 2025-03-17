package main

import "fmt"

func main(){
	var angka [5]int

	angka[0] = 1
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50

	fmt.Println("isi aray angka :", angka[1])

	for i,v:= range angka{
		fmt.Println(i,v)
	}

} 
