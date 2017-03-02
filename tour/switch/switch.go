package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")

    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("McOS.")
    case "linux":
        fmt.Println("Loonix.")
    default:
        fmt.Println("%s.", os)
    }
}
