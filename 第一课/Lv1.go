package main

import (
"fmt"
)

func main(){
	x:=-1.5
	y:=1.5
	var z float64
	for y=1.5; y>-1.5; y-=0.1{
		for x=-1.5; x<1.5; x+=0.05{
			z=x*x+y*y-1.0
			if z*z*z-x*x*y*y*y<=0{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
