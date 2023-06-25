package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minIntSlice(array []int) (min int) {
    m := 0
    index := 0
    for i, e := range array {
        if i==0 || e < m {
            m = e
            index = i
        }
    }
    return index
}

func main () {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Printf("Failed to open the CSV file: %s", "input.csv")
        panic(0)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    temp := 0
    var topthree [3]int

    for {
        line, err := reader.ReadString('\n')
        if err != nil { break }
        if line == "\n" {
            min := minIntSlice(topthree[:])

            if temp > topthree[min]{
                topthree[min] = temp
            }
            temp = 0
        }
        calories, err := strconv.Atoi(strings.TrimSpace(line))
        temp += calories
    }

    result := 0
    for _, x := range topthree {
        result += x
    }
    fmt.Println(result)

}
