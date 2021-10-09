package main

//import "fmt"
//import "time"

import (
	"fmt"
	"time"
)

func main() { //函数的“{”与函数名要求在同一行
	time.Sleep(time.Second * 1)
	// 表达式可不加 “;”,建议不加
	fmt.Println("Hello Word!")
}
