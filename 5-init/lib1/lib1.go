package lib1 // Package lib1 package 名必须与所在包名一致

import "fmt"

// Foo lib1对外提供的接口
func Foo() { // 公有函数的名字以大写字母开头
	fmt.Println("use lib1 func foo")
}

func init() {
	fmt.Println("I'm lib1 init()")
}
