

package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("  ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print(k, " ")
		}

		fmt.Println()
	}
}
