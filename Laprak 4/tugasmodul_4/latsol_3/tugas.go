package main

import "fmt"

func main() {
	var harga, diskon int

	fmt.Scan(&harga)
	fmt.Scan(&diskon)

	totalharga := harga - ((diskon * harga) / 100)

	fmt.Print("total harga nya adalah: ", totalharga)

}
