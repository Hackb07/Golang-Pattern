package main

import "fmt"

func main(){
	l := 5
	for i := 0 ; i <= l ;i++{
		for m := 0 ; m <= l-i ; m++{
			fmt.Print(" ")
		}
		for n :=1 ; n <= (2*i-1) ; n++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := l-1 ; i >= 0 ; i--{
		for m := 0 ; m <= l-i ; m++{
			fmt.Print(" ")
		}
		for n :=1 ; n <= (2*i-1) ; n++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}