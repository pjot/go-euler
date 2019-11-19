package main

import "fmt"
import "strconv"

func reverse(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return string(b)
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	return s == reverse(s)
}

func main() {
	largest := 0
	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			n := a * b
			if isPalindrome(n) && n > largest {
				largest = n
			}
		}
	}
	fmt.Println("largest", largest)
}
