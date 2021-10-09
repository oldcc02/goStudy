package main

import (
	//_ "gostudy/5-init/lib1" // 匿名导入,会调用对应包中的init函数
	//mylib "gostudy/5-init/lib2"
	. "gostudy/5-init/lib2"
)

func main() {
	Bar()
}

//import (
//	"gostudy/5-init/lib1" // 路径相对于系统环境变量中设置的GOPATH
//	"gostudy/5-init/lib2"
//)
//
//func main() {
//	lib1.Foo() // 如果不使用lib1 需要在import中注释掉lib1的导入
//	lib2.Bar()
//}
