package main

import (
	"handson2/soaldua"
	"handson2/soalempat"
	"handson2/soallima"
	"handson2/soalsatu"
	"handson2/soaltiga"
)

func main() {
	//soal 1, array GO, C
	soalsatu.SoalSatu()
	println("=======End of Soal 1=======")

	//soal 2, slice prima
	soaldua.SoalDua()
	println("=======End of Soal 2=======")

	//soal 3, slice with make
	soaltiga.SoalTiga()
	println("=======End of Soal 3=======")

	//soal 4, slice and for range
	soalempat.SoalEmpat()
	println("=======End of Soal 4=======")

	//soal 5, map
	soallima.SoalLima()
	println("=======End of Soal 5=======")
}
