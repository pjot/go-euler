package main

import "fmt"
import "../common"

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

func main() {
    for n := range triangles() {
        if len(common.Divisors(n)) > 500 {
            fmt.Println("found", n)
            break
        }
    }
}
