package main

import "fmt"

type santri struct{
	name string
	jurusan string
	asal string
	angkatan int8
}

func (s santri)InfoList(){
	fmt.Printf("name: s% \n jurusan: s% \n asal: s% \n angkatan: d% \n", s.name , s.jurusan, s.asal,s.angkatan)
}
func main(){
	santri := santri{ "irul","programmer", " 12411",  20}
	santri.InfoList()
}