package main

import "fmt"

func main() {
	var rupiah, dollar float64
	fmt.Print("MAsukkan nominal rupiah: ")
	fmt.Scan(&rupiah)
	dollar = (rupiah / 15000)
	fmt.Print("jadi ", rupiah, " rupiah = ", dollar, "dollar ")
}
