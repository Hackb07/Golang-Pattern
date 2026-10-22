package main

import "fmt"

func main(){
    n := 5
    for i:= 0 ; i <= n ; i++ {
        for j := 0 ; j <= n ; j++ {
            if i == 0 || i == n || j == 0 || j == n {
                fmt.Print("*")
              }else{
                  fmt.Print(" ")
                }
          }
		  fmt.Println()
      }
  }
