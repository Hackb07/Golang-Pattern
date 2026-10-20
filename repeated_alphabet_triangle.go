
package main

import "fmt"

func main() {
	for i := 0; i<= 5; i++{
		ch := 'A'+i
		for j := 0 ; j <= i ; j++{
			fmt.Printf("%c ",ch)
		}
		ch++
		fmt.Println()
	}
}
