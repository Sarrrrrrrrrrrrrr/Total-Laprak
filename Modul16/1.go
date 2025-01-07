package main

import "fmt"

func main() {
	var total, bilangan float64
	var count int

	for {
		fmt.Scan(&bilangan)

		if bilangan == 9999 {
			break
		}
		total += bilangan
		count++
	}
	if count == 0 {
		fmt.Println("Tidak ada bilangan yang valid untuk dihitung rata-rata.")
	} else {
		rerata := total / float64(count)
		fmt.Printf("Rata-rata dari bilangan yang dimasukkan adalah: %f\n", rerata)
	}
}
