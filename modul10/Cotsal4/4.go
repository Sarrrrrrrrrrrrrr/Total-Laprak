package main

import "fmt"

func main() {
	var b, d1, d2, d3, d4 int
	var teks string
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	d4 = b % 10
	d3 = (b % 100) / 10
	d2 = (b % 1000) / 100
	d1 = b / 1000

	if d1 < d2 && d2 < d3 && d3 < d4 {
		teks = "terurut membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		teks = "terurut mengecil"
	} else {
		teks = "tidak terurut"
	}
	fmt.Println("Digit pada bilangan", b, teks)
}
