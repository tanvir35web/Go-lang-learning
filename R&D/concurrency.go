package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg)
        time.Sleep(time.Millisecond * 500)  // Simulate work with sleep
    }
}

func main() {
    go printMessage("Hello from goroutine")  // Start a goroutine
    printMessage("Hello from main")          // Run in main thread
}
