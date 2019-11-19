package common

import "sort"

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

func Sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}
