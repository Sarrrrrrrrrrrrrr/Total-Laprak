package main

import "fmt"

func main() {
	var orang, totalmotor int
	fmt.Scan(&orang)
	totalmotor = orang / 2
	if orang%2 != 0 {
		totalmotor += 1
	}
	fmt.Println(totalmotor)
}
