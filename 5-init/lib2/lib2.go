package lib2

import "fmt"

// Bar lib2对外提供的接口
func Bar() {
	fmt.Println("use lib2 func bar")
}

func init() {
	fmt.Println("I'm lib2 init()")
}
