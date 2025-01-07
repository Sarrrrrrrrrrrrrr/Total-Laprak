package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	teks := "negatif"
	if bilangan >= 0 {
		teks = "positif"
	}
	fmt.Println(teks)
}
