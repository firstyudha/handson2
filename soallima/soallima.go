package soallima

import "fmt"

func SoalLima() {
	m := make(map[string]int)

	// beri nilai "Answer" dengan 42
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// ganti nilai "Answer" dengan 48
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// hapus "Answer
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// gunakan pengecekan: elem, ok = m[key]
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
