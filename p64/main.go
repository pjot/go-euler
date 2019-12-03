package main

import "fmt"
import "math"
import "../common"

func hash (a, d, m int) string {
    return fmt.Sprintf("%d-%d-%d", a, d, m)
}

func continuedFraction (s int) int {
    as := []string{}
    m := 0
    d := 1
    sqrtS := math.Sqrt(float64(s))
    a := int(math.Floor(sqrtS))
    for {
        m = d * a - m
        d = (s - m * m) / d
        if d == 0 {
            return 0
        }
        a = int(math.Floor((sqrtS + float64(m)) / float64(d)))
        index := common.Index(as, hash(a, d, m))
        if index >= 0 {
            return len(as) - index
        }
        as = append(as, hash(a, d, m))
    }
}

func main () {
    oddPeriods := 0
    for i := 2; i < 10001; i++ {
        period := continuedFraction(i)
        if period > 0 && period % 2 == 1 {
            oddPeriods++
        }
    }
    fmt.Println("odd periods", oddPeriods)
}
