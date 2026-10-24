
package main

import "fmt"

func printRows(level int, n int) {
    for i := 0; i < n-level; i++ {
        fmt.Print("  ")
    }

    for j := 0; j < 2*level-1; j++ {
        if j%2 == 0 {
            fmt.Print("1")
        } else {
            fmt.Print("0")
        }

        if j < 2*level-2 {
            fmt.Print(" ")
        }
    }

    fmt.Println()
}

func main() {
    n := 4

    // Upper half
    for level := 1; level <= n; level++ {
        printRows(level, n)
    }

    // Lower half
    for level := n - 1; level >= 1; level-- {
        printRows(level, n)
    }
}
