package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for m := 1 ; m <= 5-i; m++ {
			fmt.Print(" ")
		}
		for j := i; j >= 1; j--{
			fmt.Print(j)
		}
		for k := 2 ; k <= i; k++ {
			fmt.Print(k)
		}

		fmt.Println()
	}
}