package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func priority(line string) (value int) {
  n := len(line)
  left := line[:n/2] 
  right := line[n/2:]
  var item rune 

  looper:
  for _, ruckl := range left {
    for _, ruckr := range right {
      if ruckl == ruckr {
        item = ruckl
        break looper 
      }
    }
  }

  x := int(rune(item))

  if x < 97{
    return int(rune(item)-38)
  } else {
    return int(rune(item)-96)
  }
}

func main() {
  file, err := os.Open("input")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  total := 0

  // Get contents of rucksack for each line
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    strings.Split(line, " ")

    total += priority(line)
  }

  fmt.Printf("Total is: %d", total)
}
