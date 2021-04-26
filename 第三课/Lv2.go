package main

import (
	"fmt"
	"os"
)

func main(){
	file,er:=os.Create("proverb.txt")
	if er!=nil{
		fmt.Println("创建失败")
		return
	}
	file.Write([]byte("don't communicate by sharing memory share memory by communicating"))

	file,err:=os.Open("proverb.txt")
	if err!=nil{
		fmt.Println("未打开文件")
		return
	}

	text:=make([]byte,999)
	n,erro:=file.Read(text[:])
	if erro!=nil{
		fmt.Println("读取内容失败")
		return
	}
	fmt.Printf("读取了%d个数据\n",n+1)
	fmt.Printf("%s", text[:n])
	file.Close()
}
