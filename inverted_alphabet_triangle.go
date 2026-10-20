

package main

import "fmt"
func main() {
	for i:= 0 ; i<= 5;i--{
		for j:= 0 ; j<=i; j++{
			ch := 'A'
			fmt.Print(" ",ch)
		}
		fmt.Println()
	}
}
