package main

import "fmt"

func main(){
	ch := 'A'
	num := 1
	for i := 0 ; i <= 5 ; i ++{
		for j := 1; j<= i ;j++{
			if (i)%2 == 0 {
				fmt.Print(" ",num)
				num++
			}else{
				fmt.Printf(" %c", ch)
				ch++
			}
		}
		fmt.Println()
	}
}