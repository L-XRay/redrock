package main

import (
	"fmt"
	"sync"
	"time"
)

var (
		myres = make(map[int]int, 20)
		mu sync.Mutex
)

func factorial(n int,ch chan bool) {
	var res = 1
	for i := 1; i <= n; i++ {
			res *= i
	}
		myres[n] = res
		ch<-true
}

func main() {
	ch:=make(chan bool)
	for i := 1; i <= 20; i++ {
		go factorial(i,ch)
	}
	for i, v := range myres {
		<-ch
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	time.Sleep(time.Second*2)
}