package common

import "sort"
import "math"

func PrimesUntil(n int) []int {
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

func Divisors(n int) []int {
    divs := ProperDivisors(n)
    divs = append(divs, n)
    return divs
}

func ProperDivisors(n int) []int {
    limit := int(math.Sqrt(float64(n)))
    divs := []int{}
    for i := 1; i < limit + 1; i++ {
        if n % i == 0 {
            divs = append(divs, i)
            if n/i != i && i != 1{
                divs = append(divs, n/i)
            }
        }
    }
    return divs
}

func Contains(ns []int, n int) bool {
    for _, v := range ns {
        if n == v {
            return true
        }
    }
    return false
}

func Index(ns []string, n string) int {
    for k, v := range ns {
        if n == v {
            return k
        }
    }
    return -1
}

func Sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}
