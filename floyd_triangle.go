package main

import "fmt"

func main() {
	num := 1
	for i:= 0 ; i<= 5; i++{
		for j := 0 ; j<= i ; j++{
			fmt.Print(" ",num)
			num ++
		}
		fmt.Println()
	}
}
