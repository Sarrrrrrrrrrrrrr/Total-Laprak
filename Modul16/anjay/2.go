package main

import "fmt"

func main() {
	var n, inside int
	var seed int = 12345

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		seed = (seed*32719 + 3) % 32749
		x := float64(seed) / 32749.0

		seed = (seed*32719 + 3) % 32749
		y := float64(seed) / 32749.0

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			inside++
		}
	}

	fmt.Printf("Topping pada Pizza: %d\n", inside)
	fmt.Printf("PI : %.10f\n", 4.0*float64(inside)/float64(n))
}
