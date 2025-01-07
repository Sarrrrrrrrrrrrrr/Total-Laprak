package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	tambah := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			tambah++
		}
	}
	fmt.Printf("terdapat %d bilangan ganjil", tambah)
}
