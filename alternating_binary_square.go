package main 

import "fmt"

func main(){
	for i := 0; i<= 5 ; i++{
		for j  := 0 ;j<= 5;  j++{
			if (i+j)%2 == 0{
				fmt.Print("1 ")
			}else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
