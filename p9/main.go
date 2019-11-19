package main

import "fmt"

func triplet() (int, int, int) {
	for a := 1; a < 998; a++ {
		for b := 1; b < 998; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}

func main() {
	a, b, c := triplet()
	fmt.Println("product", a*b*c)
}
