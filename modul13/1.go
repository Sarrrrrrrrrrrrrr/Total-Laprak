package main

import "fmt"

func main() {
	var n int
	var p string
	fmt.Scan(&n, &p)
	c := 0
	for done := false; !done; {
		fmt.Println(p)
		c++
		done = (c >= n)
	}
}
