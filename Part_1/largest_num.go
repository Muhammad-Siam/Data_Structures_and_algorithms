package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && a > c {
		fmt.Println(a)
	} else if b > a && b > c {
		fmt.Println(b)
	} else if c > a && c > b {
		fmt.Println(c)
	}
}
