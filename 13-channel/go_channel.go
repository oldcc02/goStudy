package main

import (
	"fmt"
	"time"
)

func add(x int, y int, c chan int) {
	time.Sleep(time.Second * 2)
	// 传值给channel channel <- value
	c <- x + y
}

func addThree(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
	// 如果不手动关闭channel 使用range会阻塞导致报错
	// range会认为协程还在写入值到channel
	close(c)
}
func main() {
	// 无缓冲创建
	c := make(chan int)
	x := 10
	y := 12
	go add(x, y, c)
	go add(x, y+1, c)
	go add(x+1, y, c)
	// channel中获取值
	// 方法一 param,ok := <-channel
	// 使用这种方法获取channel中的值会发生阻塞
	// 直到等到channel中有值传过来
	z1, ok := <-c
	if ok {
		fmt.Println(z1)
	}

	z2, ok := <-c
	if ok {
		fmt.Println(z2)
	}

	z3, ok := <-c
	if ok {
		fmt.Println(z3)
	}

	// 方法二 for item := range channel
	c2 := make(chan int)
	go addThree(c2)
	for item := range c2 {
		fmt.Println(item)
	}
}
