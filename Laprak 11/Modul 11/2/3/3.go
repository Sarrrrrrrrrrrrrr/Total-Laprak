package main

import "fmt"

func main() {
	var bilangan, hasil int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	switch {
	case bilangan%2 != 0: // Bilangan Ganjil
		if bilangan%5 == 0 && bilangan > 5 {
			hasil = bilangan * bilangan
			fmt.Println("Bilangan Kelipatan 5")
			fmt.Printf("Hasil kuadrat dari %d^2 adalah %d\n", bilangan, hasil)
		} else {
			hasil = bilangan*2 + 1
			fmt.Println("Bilangan Ganjil")
			fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
		}
	case bilangan%2 == 0: // Bilangan Genap
		if bilangan%10 == 0 && bilangan > 10 {
			hasil = bilangan / 10
			fmt.Println("Bilangan Kelipatan 10")
			fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
		} else {
			hasil = bilangan * (bilangan + 1)
			fmt.Println("Bilangan Genap")
			fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
		}
	default:
		fmt.Println("Masukkan angka yang valid.")
	}
}
