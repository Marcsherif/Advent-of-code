package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func priority(arr []string) (value int) {
  var item string

  looper:
  for _, x := range str1 {
    for _, j := range str2 {
      for _, k := range str3 {
        if x == j && j == k{
          item = x 
          fmt.Println(item)
          break looper 
        }
      }
    }
  }
  return 30
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
  count := 1
  var fline []string

  // Get contents of rucksack for each line
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    strings.Split(line, " ")

    // Call Prio every 3 times the for loop excecutes
    if count % 3 == 0 { 
        fline = append(fline, line)
        total += priority(fline) 
        fline = nil
    } else {
        fline = append(fline, line)
        //fmt.Println(fline)
    }
    count++
  }

  fmt.Printf("Total is: %d", total)
}
