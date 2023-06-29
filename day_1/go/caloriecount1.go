package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
    file, err := os.Open("./go/input.txt")

    if err != nil {
        fmt.Printf("Failed to open the txt file: %s", "input.txt")
        panic(0)
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    temp := 0
    max := 0

    for {
        line, err := reader.ReadString('\n')
        if err != nil { break }
        if line == "\n" {
            if temp > max{
                max = temp
            }
            temp = 0
        }
        calories, err := strconv.Atoi(strings.TrimSpace(line))
        temp += calories
    }

    fmt.Println(max)

}

