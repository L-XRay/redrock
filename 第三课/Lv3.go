package main

import (
	"fmt"
)

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {       //协程内的i与循环中的i为同一个，当循环跑得足够快时，协程已经多次取了一个值
			fmt.Println(i)
			if i == 9 {
				over <- true
			}
		}(i)
		//if i == 9 {          //管道在同一协程内无法传递
		//	over <- true
		//}
	}
	<-over
	fmt.Println("over!!!")
}