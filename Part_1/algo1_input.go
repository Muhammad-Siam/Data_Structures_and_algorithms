package main

import "fmt"

func main() {
	fmt.Println("Enter your age:")
	var n1 int
	fmt.Scanln(&n1)
	var n2 int
	fmt.Scanln(&n2)
	sum:= n1+n2
	fmt.Println(sum)
}