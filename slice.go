package main

import "fmt"

func main (){
	name := [...]string{"eko" , "kurniawan", "irul", "fauzan", "fathoni", "huda"}

	slice1 := name[4:6]
	fmt.Println(slice1)

	slice2 := name[:3]
	fmt.Println(slice2)

	slice3 := name[3:] 
	fmt.Println(slice3)

	slice4 := name[:]
	fmt.Println(slice4)

	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] //sabtu dan minguu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru" 
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "sabtu lama"
	//	days baru := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2,5)
	newSlice[0] = "irul"
	newSlice[1] = "iruljg"
	// newSlice[2] = "irul2" //error malahan  dianya karna kapasitasnya sudah di deklarasikan yaitu 2
	//jika ingin menambahkanya lebihbaik menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "irulSlice2")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "irulBaru"
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	formSlice := days[:]
	toSlice := make([]string, len(formSlice), cap(formSlice))

	copy(toSlice, formSlice)

	fmt.Println(formSlice)
	fmt.Println(toSlice)

	iniaray := [...]int{1,2,3,4,5,6,7} //bedanya di kurung iku di beri value/penjelasan .../angka
	iniSlace := []int{1,2,3,4,5,6,7} //jika ingin membuat slice tinggal value di dalam kurung siku tingal di kosongi saja

	fmt.Println(iniaray)
	fmt.Println(iniSlace)


}