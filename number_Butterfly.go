
package main

import "fmt"

func printRow(level, n int) {
    // Left-side numbers
    for i := 1; i <= level; i++ {
        fmt.Print(i)
    }

    // Middle spaces
    if level < n {
        for i := 0; i < 2*(n-level)-1; i++ {
            fmt.Print(" ")
        }
    }

    // Right-side numbers
    for i := level; i >= 1; i-- {
        fmt.Print(i)
    }

    fmt.Println()
}

func main() {
    n := 4

    for level := 1; level <= n; level++ {
        printRow(level, n)
    }

    for level := n - 1; level >= 1; level-- {
        printRow(level, n)
    }
}
