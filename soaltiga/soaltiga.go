package soaltiga

import "fmt"

func SoalTiga() {

	// Buat slice of int bernama "a" dengan panjang 5
	a := make([]int, 5)
	printSlice("a", a)

	// Buat slice of int bernama "b" dengan panjang 0 kapasitas 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	// Buat variabel c dengan dua data awal b
	c := b[:2]
	printSlice("c", c)

	// Buat variabel d dengan data b index ke 2 sampai 4
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
