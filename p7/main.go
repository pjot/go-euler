package main

import "fmt"
import "sort"
import "math"

func primesUntil(n int) []int {
	res := make(map[int]bool)
	i := 2
	for i < n {
		v, ok := res[i]
		if !ok {
			res[i] = true
		}
		if !ok || v {
			for j := i * i; j < n; j += i {
				res[j] = false
			}
		}
		i++
	}

	keys := []int{}
	for k, v := range res {
		if v {
			keys = append(keys, k)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func lg(f float64) float64 {
	return math.Log(f)
}

func pi(num int) int {
	n := float64(num)
	return int(n * (lg(n) + lg(lg(n))))
}

func main() {
	primes := primesUntil(pi(10001))
	fmt.Println("prime", primes[10000])
}
