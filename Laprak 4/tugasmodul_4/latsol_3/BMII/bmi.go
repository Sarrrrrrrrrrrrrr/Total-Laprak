package main

import "fmt"

func main() {
	var bmi, tb, bb float64
	fmt.Println("masukkan bmi")
	fmt.Scan(&bmi)

	fmt.Println("masukkan bb")
	fmt.Scan(&tb)

	bb = bmi * (tb * tb)

	fmt.Printf("berat badan anda adalah %.f", bb)
}
