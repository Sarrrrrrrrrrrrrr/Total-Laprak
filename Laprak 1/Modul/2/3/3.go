package main

import "fmt"

func main() {
	var n, s1, s2, j, temp int
	fmt.Scan(&n)
	s1 = 0
	s2 = 1
	j = 0
	for j < n {
		fmt.Print(s1, " ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
		j = j + 1
	}
}
