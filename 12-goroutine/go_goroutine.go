package main

import (
	"fmt"
	"runtime"
	"time"
)

func task() {
	for true {
		fmt.Println("子协程")
		time.Sleep(time.Second)
		// 退出协程
		runtime.Goexit()
	}
}

func main() {
	// 使用go关键字创建goroutine
	go task()

	for true {
		fmt.Println("主协程")
		time.Sleep(time.Second)
	}
}
