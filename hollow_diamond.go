package main

import "fmt"

func main() {
    n := 5

    // Upper hollow pyramid
    for i := 1; i <= n; i++ {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }

        for k := 1; k <= 2*i-1; k++ {
            if k == 1 || k == 2*i-1 {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }

        fmt.Println()
    }

    // Lower hollow inverted pyramid
    for i := n - 1; i >= 1; i-- {
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }

        for k := 1; k <= 2*i-1; k++ {
            if k == 1 || k == 2*i-1 {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }

        fmt.Println()
    }
}
