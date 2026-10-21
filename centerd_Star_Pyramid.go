package main

import "fmt"

func main(){
	for i := 0 ; i <= 7 ;i++{
		for m := 0 ; m <= 7-i ; m++{
			fmt.Print(" ")
		}
		for n :=1 ; n <= (2*i-1) ; n++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}