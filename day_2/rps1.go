package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {
    file, err := os.Open("day_2/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    score := 0 

    // Get my and enemy's shape for each line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        strings.Split(line, "")
        enemy := line[0]
        shape := line[2]
        // fmt.Printf("enemy: %q me: %q\n", enemy, shape) // use to see each round

        // Count the shape and check outcome of round
        switch shape {
            case 'X': // rock
            score += 1
            if enemy == 'A' { score += 3}
            if enemy == 'B' { score += 0}
            if enemy == 'C' { score += 6}
            case 'Y': // paper
            score += 2 
            if enemy == 'A' { score += 6}
            if enemy == 'B' { score += 3}
            if enemy == 'C' { score += 0}
            case 'Z': // scissors
            score += 3
            if enemy == 'A' { score += 0}
            if enemy == 'B' { score += 6}
            if enemy == 'C' { score += 3}
        }
    }
    fmt.Println("My total score according to the strategy guide is:",score)

}
