package main

import (
	"fmt"
)

func main() {
	var r float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)
	pi := 3.1415926535
	v := (4.0 / 3.0) * pi * (r * r * r)
	l := 4 * pi * (r * r)

	fmt.Println("Bola dengan jari-jari", r, "memiliki volume", v, "dan luas permukaan", l)
}
