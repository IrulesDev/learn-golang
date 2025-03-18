package main

import "fmt"

func main(){

	type noKtp string

	var irulKTP noKtp = "11111111"

	var contoh string = "22222222"
	var  contohKtp noKtp = noKtp(contoh)

	fmt.Println(irulKTP)
	// fmt.Println(contoh)
	fmt.Println(contohKtp)
}