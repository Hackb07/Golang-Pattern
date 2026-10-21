package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		for j := 0; j <= i-1; j++ {
			fmt.Print(" ", i+j)
		}
		fmt.Println()
	}
}
