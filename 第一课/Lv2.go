package main

import (
	"fmt"
	"strconv"
)
func sum(a int,b int)int{
	return a+b
}
func sub(a int,b int)int{
	return a-b
}
func mul(a int,b int)int{
	return a*b
}
func div(a int,b int)int{
	return a/b
}
func main() {
	var n [3]string
	var i int
	a:=n[0:1]//var a,b int
	c:=n[1:2]//var c byte
	b:=n[2:3]
	for {
		fmt.Println("input:")
		for i = 0; i < 3; i++ {
			fmt.Scan(&n[i]) //可直接使用  fmt.Scanf("%d%c%d", &a, &c, &b)
		}
		n1, err := strconv.Atoi(a[0]) //n1存储第一个数字字符转换后的数值
		if err != nil {
			fmt.Println(err)
		}
		n2, err := strconv.Atoi(b[0]) //n2存储第二个数字字符转换后的数值
		if err != nil {
			fmt.Println(err)
		}
		if c[0] == "+" {
			fmt.Println("output:")
			fmt.Println(sum(n1, n2))
		} else if c[0] == "-" {
			fmt.Println("output:")
			fmt.Println(sub(n1, n2))
		} else if c[0] == "*" {
			fmt.Println("output:")
			fmt.Println(mul(n1, n2))
		} else if c[0] == "/" {
			fmt.Println("output:")
			fmt.Println(div(n1, n2))
		}
	}
}

