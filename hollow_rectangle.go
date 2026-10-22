package main

import "fmt"

func main(){
	l := 6
	b := 3
	for i := 1 ; i <= l ; i ++ {
		for j := 1; j <= b ; j++ {
			if i == 1 || i == l || j == 1 || j == b {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
