

package main

import "fmt"

func main() {
    n := 5

    for i := 1; i <= n; i++ {
        // Print leading spaces
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }

        // Print pyramid pattern
        for k := 1; k <= 2*i-1; k++ {
            if i == n || k == 1 || k == 2*i-1 {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }

        fmt.Println()
    }
}
