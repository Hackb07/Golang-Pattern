package main

import "fmt"

func main() {
	n := 5
	ch := 'A'

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if j%2 == 1 {
				fmt.Printf("%c", ch)
				ch++
			} else {
				fmt.Print(j - 1)
			}
		}
		fmt.Println()
	}
}
