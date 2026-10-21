package main

import "fmt"

func main(){
	for i := 5 ; i >= 0 ; i--{
		for m := 0 ; m <= 5-i ; m++{
			fmt.Print(" ")
		}
		for n :=1 ; n <= (2*i-1) ; n++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}