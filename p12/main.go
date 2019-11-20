package main

import "fmt"
import "math"

func triangles() <-chan int {
	out := make(chan int)
	go func() {
        n := 0
		for i := 1; ; i += 1 {
            n += i
			out <- n
		}
		close(out)
	}()
	return out
}

func divisors(n int) []int {
    limit := int(math.Sqrt(float64(n)))
    divs := []int{}
    for i := 1; i < limit + 1; i++ {
        if n % i == 0 {
            divs = append(divs, i)
            if n /i != i {
                divs = append(divs, n/i)
            }
        }
    }
    return divs
}

func main() {
    for n := range triangles() {
        if len(divisors(n)) > 500 {
            fmt.Println("found", n)
            break
        }
    }
}
