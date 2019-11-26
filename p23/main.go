package main

import "fmt"
import "../common"

func main() {
    abundant := []int{}
    limit := 28124
    for i:= 1;i < limit;i++{
        divisors := common.ProperDivisors(i)
        divisorSum := common.Sum(divisors)
        if divisorSum > i {
            abundant = append(abundant, i)
        }
    }

    seenNumbers := make(map[int]bool)
    for i:=0; i < limit;i++ {
        seenNumbers[i] = false
    }
    for i:=0;i < len(abundant); i++ {
        for j := i; j < len(abundant); j++ {
            sum := abundant[i] + abundant[j]
            seenNumbers[sum] = true
        }
    }
    finalSum := 0
    for i, seen := range seenNumbers {
        if !seen {
            finalSum += i
        }
    }
    fmt.Println("sum", finalSum)
}
