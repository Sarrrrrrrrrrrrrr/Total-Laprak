package main

import "fmt"

func main() {
	var a, b, hasil float64
	fmt.Scan(&a, &b)
	if b != 0 {
		hasil = a / b
		fmt.Print("hasil pembagian adalah", hasil)
	} else {
		fmt.Print("variabel b bernilai nol")
	}
	fmt.Println("Program Selesai")
}
