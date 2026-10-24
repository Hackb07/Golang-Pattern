package main

import "fmt"

func main() {
	n := 5

	// Upper part
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*(n-i)+1; k++ {
			if k == 1 || k == 2*(n-i)+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	// Lower part
	for i := n - 1; i >= 1; i-- {
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*(n-i)+1; k++ {
			if k == 1 || k == 2*(n-i)+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
