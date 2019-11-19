package main

import "fmt"

func main() {
	limit := 1000
	i := 1
	sum := 0
	for i < limit {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i++
	}
	fmt.Println("sum", sum)
}
