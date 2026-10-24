package main

import "fmt"

func main(){
	rows := 30
	columns := 100
	for i:= 0 ; i <= rows ; i++ {
		for j := 0 ; j <= columns ; j++ {
			if (i == 0 && j%4 == 0) || (i == 1 && j%2 == 1) || (i == 2 && j%4 == 2){
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
