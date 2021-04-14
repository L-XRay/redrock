package main

import (
	"fmt"
)

func f(n[10]int,size int)int{
	if n[0]<=0{
		n[0]=-n[0]
	}
	min:=n[0]
	var i int
	for i = 1; i < size; i++ {
		if n[i]<=0{
			n[i]=-n[i]
		}
		if n[i]<min{
			min=n[i]
		}
	}
	return min
}
func main() {
	var n [10]int
	var i int
	fmt.Println("input:")
	for i = 0; i < 10; i++ {
		fmt.Scan(&n[i])
	}

	fmt.Println("output:")
	for i = 0; i < 10; i++ {
		fmt.Print(n[i]," ")
	}
	fmt.Println()
	fmt.Printf("min=%d",f(n,i))
}


