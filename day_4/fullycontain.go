package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fullycontain(splitted []int) (int) {
    contains := 0
    // Check first index of both elf's assignment
    if splitted[0] <= splitted[2] {
        if splitted[1] >= splitted[3] {
           contains = 1 
        }
    }

    if splitted[0] >= splitted[2] {
        if splitted[1] <= splitted[3] {
           contains = 1 
        }
    }

    return contains
}

func overlap(splitted []int) (int) {
    contains := 0

    // EX: 2-6, 4-8
    // Check first index of both elf's assignment
    if splitted[0] <= splitted[2] {     // 2 < 4
        // Check second index of first elf and first of second elf
        if splitted[1] >= splitted[2] { // 6 > 4
           contains = 1 
        }
    }

    // EX: 4-8, 2-6
    if splitted[0] >= splitted[2] {     // 4 > 2
        if splitted[0] <= splitted[3] { // 4 < 6
           contains = 1 
        }
    }

    return contains
}

func convtoints(splitted []string)([]int) {

    var ints []int
    ints = make([]int, 3)
    for i := 0; i<=1; i++ {
        for j:=0; j<=1; j++ {
            extrasplit := strings.Split(splitted[i], "-")
            intelegi, err := strconv.Atoi(extrasplit[j])
            _ = err
            ints = append(ints, intelegi)
        }
    }
    return ints
}

func main () {
    file, err := os.Open("./input")

    if err != nil {
        fmt.Printf("Failed to open the txt file: %s", "input")
        panic(0)
    }
    defer file.Close()

    fcounter := 0
    ocounter := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        splitted := strings.Split(line, ",")

        ints := convtoints(splitted)        // convert string to int array
        fcontains := fullycontain(ints[3:]) // check for fully containing tasks
        ocontains := overlap(ints[3:])      // check for overlapping tasks
        fcounter += fcontains               // update counter if overlapping tasks
        ocounter += ocontains
    }

    fmt.Println("Fully Contains:", fcounter)
    fmt.Printf("Overlaps:%10d\n", ocounter)
}

