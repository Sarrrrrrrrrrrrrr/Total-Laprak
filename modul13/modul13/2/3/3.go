package main

import "fmt"

func main() {
	var target, donation, totalDonasi, jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for totalDonasi < target {
		jumlahDonatur++
		fmt.Printf("Masukkan donasi dari donatur %d: ", jumlahDonatur)
		fmt.Scan(&donation)
		totalDonasi += donation
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donation, totalDonasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
