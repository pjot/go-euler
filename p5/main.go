package main

import "fmt"

func dividesNicely(n int, limit int) bool {
	for i := 2; i <= limit; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 1
	for {
		if dividesNicely(i, 20) {
			fmt.Println("smallest number", i)
			break
		}
		i++
	}
}
