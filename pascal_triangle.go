package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		val := 1
		for j := 1; j <= i; j++ {
			fmt.Print(" ", val)
			if j < i {
				val = val * (i - j) / j
			}
		}
		fmt.Println()
	}
}
