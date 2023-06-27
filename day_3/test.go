package main

import "fmt"

func main() {
    fmt.Printf("%d\n", int(rune('A')-38))
    fmt.Printf("%d\n", rune('Z')-38)
    fmt.Printf("%d\n", rune('a')-96)
    fmt.Printf("%d\n", rune('z')-96)
}
