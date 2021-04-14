package main

import (
	"fmt"
)
func Receivers(v interface{})  {
	switch v.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown Type")
	}
}
func main() {
	Receivers(66.6)
	Receivers(false)
	Receivers("你好")
	Receivers(15)
}

