package main

import "fmt"
import "math"
import "sort"

func largestFactor(n int) int {
	primes := possibleFactorsFor(int(n / 2))
	for _, p := range primes {
		if n%p == 0 {
			return p
		}
	}
	return 1
}

func possibleFactorsFor(n int) []int {
	res := make(map[int]bool)
	i := 2
	limit := int(math.Sqrt(float64(n)))
	for i < limit {
		v, ok := res[i]
		if !ok {
			res[i] = true
		}
		if !ok || v {
			for j := i * i; j < limit; j += i {
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
		return keys[i] > keys[j]
	})
	return keys
}

func main() {
	fmt.Println("largest factor", largestFactor(600851475143))
}
