package main

import "fmt"

//函数格式
//func 函数名(参数1 类型, ***) (返回值 返回值类型, ***) {
//	return 返回值
//}

// 单返回值函数
func foo1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c
}

// 返回多个匿名返回值
func foo2(a string, b int) (int, int) {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c, c + 1
}

// 返回多个带名返回值
func foo3(a int, b int) (r1 int, r2 int) {

	//return a, b
	// 与下方的返回方式等价

	// 未赋值前r1,r2都为0，且作用域为函数体内
	fmt.Println(r1, r2) // ==> 0 0
	r1 = a
	r2 = b
	return
}

func main() {
	fmt.Println("----foo1----")
	c := foo1("bar", 123)
	fmt.Println(c)

	fmt.Println("----foo2----")
	r1, r2 := foo2("func2", 333)
	fmt.Println(r1, r2)

	fmt.Println("----foo3----")
	r1, r2 = foo3(2, 3)
	fmt.Println(r1, r2)
}
