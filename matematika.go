package main

import "fmt"

func main(){
	var a = 10
	var b = 20
	var c = 30
	var g = 5
	d := a + b * c / g
	fmt.Println(d)

	i := 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	j := 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	//di turunkan
	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

}