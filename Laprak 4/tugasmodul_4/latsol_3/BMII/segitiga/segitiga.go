package main

import (
	"fmt"
	"math"
)

func main() {

	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	d1 := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
	d2 := (x3-x2)*(x3-x2) + (y3-y2)*(y3-y2)
	d3 := (x1-x3)*(x1-x3) + (y1-y3)*(y1-y3)

	ab := math.Sqrt(d1)
	bc := math.Sqrt(d2)
	ca := math.Sqrt(d3)

	max := math.Max(ab, math.Max(bc, ca))
	fmt.Printf("%.2f\n", max)
}
