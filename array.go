package main

import "fmt"

func main(){
	var angka [5]int

	angka[0] = 1
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50

	fmt.Println(angka[0])
	fmt.Println(angka[1])
	fmt.Println(angka[2])
	fmt.Println(angka[3])
	fmt.Println(angka[4])

	var value = [3] int{
		60,
		50, //<-

	}

	fmt.Println(value)
} 
