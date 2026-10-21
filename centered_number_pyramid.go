package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for m := 0; m <= 5-i; m++ {
			fmt.Print(" ")
		}
		for n := 1; n <= i; n++ {
			fmt.Print(" ", n)
		}
		fmt.Println()
	}
}
