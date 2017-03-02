package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("How far is Saturday?")
    today := time.Now().Weekday()

    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("Two days hence.")
    default:
        fmt.Println("Too far away.")
    }
}
