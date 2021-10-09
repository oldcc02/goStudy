package main

import "fmt"

type book struct {
	name string
}

func test(arg interface{}) {
	fmt.Printf("arg type is %T \n", arg)
	// 类型断言
	value, flag := arg.(string)
	// 如果arg 是string类型 会执行下面代码
	if flag {
		fmt.Println("arg type is string,value = ", value)

	}
}

func main() {
	book := book{name: "哈利波特"}
	test(book)

	s := "ttt"
	test(s)
}
