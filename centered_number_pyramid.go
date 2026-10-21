package main

import "fmt"

func main() {
	const rows = 5

	// The widest row holds 2*rows-1 numbers separated by single spaces.
	maxWidth := 2*(2*rows-1) - 1

	for i := 1; i <= rows; i++ {
		count := 2*i - 1
		width := 2*count - 1

		// Left padding so every row is centred on the same axis.
		for m := 0; m < (maxWidth-width)/2; m++ {
			fmt.Print(" ")
		}
		for n := 1; n <= count; n++ {
			if n > 1 {
				fmt.Print(" ")
			}
			fmt.Print(n)
		}
		fmt.Println()
	}
}
