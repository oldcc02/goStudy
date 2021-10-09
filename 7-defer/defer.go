package main

import "fmt"

func main() {
	// defer示例
	//defer fmt.Println("go end")
	//fmt.Println("fist go")
	//fmt.Println("second go")

	// defer执行顺序
	defer func1()
	defer func2()
	defer func3()

	// defer 与 return执行顺序
	fmt.Println("defer and return")
	func4()

	fmt.Println("start defer func")

}

func func3() {
	fmt.Println("3")
}

func func2() {
	fmt.Println("2")
}

func func1() {
	fmt.Println("1")
}

func func4() int {
	defer fmt.Println("I'm defer")
	return func5()
}

func func5() int {
	fmt.Println("I'm return")
	return 0
}
