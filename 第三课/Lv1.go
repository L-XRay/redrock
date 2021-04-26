
package main

import (
	"fmt"
)

func ch1(ch chan bool,quite1 chan bool) {

	for i := 0; i <= 100; i++ {
		ch <- true
		if i%2 == 0 {
			fmt.Println("1:", i)
		}
	}
	quite1<-true
}

func ch2(ch chan bool,quite2 chan bool) {

	for i := 0; i <= 100; i++ {
		<-ch
		if i%2 == 1 {
			fmt.Println("2:", i)
		}
	}
	quite2<-true
}

func main() {
	ch := make(chan bool)
	quite1 := make(chan bool)
	quite2 := make(chan bool)
	go ch1(ch,quite1)
	go ch2(ch,quite2)
	<-quite1
	<-quite2
}