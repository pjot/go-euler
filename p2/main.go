package main

import "fmt"

func main() {
	sum := 0
	curr := 1
	next := 2
	for curr < 4000000 {
		if curr%2 == 0 {
			sum += curr
		}
		curr, next = next, curr+next
	}
	fmt.Println("sum", sum)
}
