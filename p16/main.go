package main

import "fmt"
import "math/big"
import "strconv"

func digitSum(n string) int {
    sum := 0
    for _, d := range n {
        i, _ := strconv.Atoi(string(d))
        sum += i
    }
    return sum
}

func bigPower(exponent int) string {
    r := big.NewInt(2)
    two := big.NewInt(2)
    for i := 1; i < exponent; i++ {
        r = r.Mul(r, two)
    }
    return r.Text(10)
}


func main() {
    n := bigPower(1000)
    fmt.Println("digit sum", digitSum(n))
}
