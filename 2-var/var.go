package main

import "fmt"

//声明全局变量时 1,2,3三种方法可以 第四种不行
//第四种方法:= 只适用于函数体内声明变量
func main() {
	//一. 声明变量时，默认值是0
	var a int
	fmt.Println("a = ", a) // ==> a = 0

	//二. 变量初始化值
	var b int = 10
	fmt.Println("b = ", b) // ==> b = 10
	var cc string = "cc"

	//三. 声明变量时可不指明数据类型，通过值自动匹配数据类型
	var c = "hello"
	fmt.Println("c = ", c)              // ==> c = hello
	fmt.Printf("type of c = %T\n", c)   // ==> type of c = string
	fmt.Printf("type of b = %T\n", b)   // ==> type of b = int
	fmt.Printf("type of cc = %T\n", cc) // ==> type of cc = string

	//四. (常用)省去var关键字，自动匹配数据类型
	d := 100
	fmt.Println("d = ", d)            // ==> d = 100
	fmt.Printf("type of d = %T\n", d) // ==> type of d = int

	// 多变量声明
	var q, w int = 1, 2
	fmt.Println("q = ", q, "w = ", w) // ==> q =  1 w =  2
	var e, r = "foo", "bar"
	fmt.Println("e = ", e, "r = ", r) // ==> e =  foo r =  bar

	// 使用括号多行声明
	var (
		qq int    = 100
		ww string = "too"
	)
	fmt.Println("qq = ", qq, "ww = ", ww) // ==> qq =  100 ww =  too
}
