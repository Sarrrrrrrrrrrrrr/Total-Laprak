package main

import "fmt"

func main() {
	var umur int
	var x bool
	fmt.Scan(&umur, &x)

	if umur >= 17 && x {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("blom bisa kocak")
	}
}
