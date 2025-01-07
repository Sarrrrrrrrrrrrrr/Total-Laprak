package main

import "fmt"

func main() {
	var n, y string
	fmt.Scan(&n, y)
	attempt := 0
	for n != "Admin" && y != "Admin" {
		fmt.Scan(&n, &y)
		attempt++
	}
	fmt.Print(attempt)
	fmt.Print("Percobaan gagal login")
}
