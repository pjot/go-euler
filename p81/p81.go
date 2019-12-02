package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

func readFromFile (fileName string) (map[string]int, int) {
    data, _ := ioutil.ReadFile(fileName)
    result := make(map[string]int)
    size := 0
    for x, row := range strings.Split(string(data), "\n") {
        if len(row) < 1 { continue }

        size++
        for y, i := range strings.Split(row, ",") {
            j, err := strconv.Atoi(i)
            if err == nil {
                result[hash(x, y)] = j
            }
        }
    }
    return result, size
}

func emptyDist (size int) map[string]int {
    dist := make(map[string]int)
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            dist[hash(x, y)] = 999999
        }
    }
    return dist
}

func emptyVisited (size int) map[string]bool {
    visited := make(map[string]bool)
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            visited[hash(x, y)] = false
        }
    }
    return visited
}

func hash (x, y int) string {
    return fmt.Sprintf("%d-%d", x, y)
}

func unhash (key string) (int, int) {
    nums := strings.Split(key, "-")
    xs, _ := strconv.Atoi(nums[0])
    ys, _ := strconv.Atoi(nums[1])
    return xs, ys
}

func minNode (visited map[string]bool, dist map[string]int) string {
    lowest := 9999999
    low_key := ""
    for key, val := range visited {
        if val { continue }

        v := dist[key]
        if lowest > v {
            low_key = key
            lowest = v
        }
    }
    return low_key
}

func hasUnvisited (visited map[string]bool) bool {
    for _, val := range visited {
        if !val { return true }
    }
    return false
}

func main () {
    matrix, size := readFromFile("p081_matrix.txt")
    dist := emptyDist(size)
    visited := emptyVisited(size)

    dist[hash(0, 0)] = matrix[hash(0, 0)]

    for hasUnvisited(visited) {
        key := minNode(visited, dist)
        x, y := unhash(key)
        visited[key] = true

        neighbours := []string{
            hash(x + 1, y),
            hash(x, y + 1),
        }
        for _, coords := range neighbours {
            d, _ := dist[key]

            m, m_err := matrix[coords]
            if !m_err { continue }

            dn, d_err := dist[coords]
            if !d_err { continue }

            a := d + m
            if a < dn {
                dist[coords] = a
            }
        }
    }
    fmt.Println("path", dist[hash(size - 1, size - 1)])
}
