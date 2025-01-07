package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	prime := n > 1
	for i := 2; i*i <= n && prime; i++ {
		prime = prime && n%i != 0
	}
	if prime {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan prima")
	}
}
