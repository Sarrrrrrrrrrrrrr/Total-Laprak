package main

import "fmt"

func main() {
	var Berat int
	fmt.Scan(&Berat)
	fmt.Print("Berat Parsel (gram): ")

	kg := Berat / 1000
	gr := Berat % 1000

	biayaKg := kg * 10000
	biayaGr := (gr / 100) * 500
	totalBiaya := biayaKg + biayaGr

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGr)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
