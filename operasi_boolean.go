package main

import "fmt"

func main(){
	var NilaiAkhir = 90
	var Absensi = 81

	var LulusNilaiAkhir bool = NilaiAkhir > 80
	var Lulusabsensi bool = Absensi > 80

	var lulus bool = LulusNilaiAkhir && Lulusabsensi

	fmt.Println(lulus)
}