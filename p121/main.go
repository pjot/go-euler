package main

import "fmt"

func main () {
    line := []int{1, 1}
    for row := 2; row <= 15; row++ {
        last := line
        line = []int{}
        for col := 0; col < row + 1; col++ {
            red, red_multiplier, blue := 0, 0, 0
            if col < row { blue = last[col] }
            if col > 0 {
                red = row
                red_multiplier = last[col - 1]
            }
            line = append(line, blue + red * red_multiplier)
        }
    }

    breakoff := len(line) / 2

    blue := 0
    total := 0
    for i, val := range line {
        total += val
        if i < breakoff { blue += val }
    }

    fmt.Println(total / blue)
}
