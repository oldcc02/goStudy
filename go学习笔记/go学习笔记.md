*环境*

*go：1.17.1*

*操作系统：windows10*

IDE：goland

### 介绍

go的优势

- 部署简单
  - [ ] 可直接编译为机器码
  - [ ] 不依赖其他库
  - [ ] 直接运行即可部署
- 静态类型语言
  - [ ] 编译时检查出隐藏的大多数问题

- 语言层面的并发
  - [ ] 天生的基因支持
  - [ ] 充分利用多核

- 强大的标准库支撑
  - [ ] runtime系统调度机制
  - [ ] 高效的GC垃圾回收
  - [ ] 丰富的标准库

- 简单易学

  - [ ] 25个关键字
  - [ ] 语法简洁（参考c），内嵌C语法支持
  - [ ] 面向对象特征（继承，多态，封装）
  - [ ] 跨平台

  

### 1.安装go

#### 1.1下载go

选择与自己操作系统对应的go安装包

`https://golang.google.cn/dl/`

下载号后按照提示安装

#### 1.2 配置系统环境变量

添加 GOROOT，GOPATH到系统环境变量

| 变量   | 值           | 解释               |
| ------ | ------------ | ------------------ |
| GOROOT | E:\go1.17    | go源码包的安装位置 |
| GOPATH | D:\goproject | 项目代码的存放目录 |

在终端使用 `go version`检测是否配置成功

GOPATH目录树

  -GOPATH

​		-src   *#项目源码*

​			-project1  

​			-project2

​		-bin

​		-pkg

### 2.使用go

#### 2.1 Hello Word

```go
package main

import "fmt"

func main() {
	// 表达式可不加 “;”,建议不加
	fmt.Println("Hello Word!")
}

```

#### 2.2 导包格式

```go
package main

//import "fmt"
//import "time"
// 上下两种导入方式等价
import (
	"fmt"
	"time"
)

func main() { //函数的“{”与函数名要求在同一行
	time.Sleep(time.Second * 1)
	fmt.Println("Hello Word!")
}
```

### 3.变量声明

```go
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
```

### 4 常量

```go
package main

func main() {
	// const关键字（只读）
	const l int = 10

	// const()定义枚举类型
	//const (
	//	a = 0
	//	b = 1
	//	c = 2
	//)
	// 与下方使用iota的效果一致

	// iota 的使用
	const (
		a = iota // a=0
		b        // b=1
		c        // c=2
	)

	const (
		aa, bb = iota + 1, iota + 2 //iota=0, aa=iota+1=1, bb=iota+2=2
		cc, dd                      //iota=1, cc=iota+1=2, dd=iota+2=3
		ee, ff = iota * 2, iota * 3 //iota=2, ee=iota*2=4, ff=iota*3=6
		gg, hh                      //iota=3, gg=iota*2=6, hh=iota*3=9
	)

	// iota只能在const()中使用 下面这种方式不可取
	//var o int = iota
}

```

###  5 函数

#### 5.1 函数以及函数返回值的格式

函数格式

```go
func 函数名(参数1 类型, ***) (返回值 返回值类型, ***) {
	return 返回值
}
```



```go
package main

import "fmt"

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

```

#### 5.2 init函数与import导包

![](D:\goproject\src\gogogo\go学习笔记\image\导包流程.png)

ps:

1. init函数在main函数前执行
2. 未使用的包需要注释掉，否则编译报错

创建类似文件树

![](C:\Users\Admin\Desktop\go学习笔记\image\导包文件树.png)

main.go

```go
package main

import (
	"gostudy/5-init/lib1" // 路径相对于系统环境变量中设置的GOPATH
	"gostudy/5-init/lib2"
)

func main() {
    lib1.Foo() // 如果不使用lib1 需要在import中注释掉lib1的导入
	lib2.Bar()
}
```

lib1.go

```go
package lib1 // package 名必须与所在包名一致

import "fmt"

// lib1对外提供的接口
func Foo() { // 公有函数的名字以大写字母开头
	fmt.Println("use lib1 func foo")
}

func init() {
	fmt.Println("I'm lib1 init()")
}
```

lib2.go

```go
package lib2

import "fmt"

// lib2对外提供的接口
func Bar() {
	fmt.Println("use lib2 func bar")
}

func init() {
	fmt.Println("I'm lib2 init()")
}
```

#### 5.3 import匿名与别名导入

```go
package main

import (
	_ "gostudy/5-init/lib1"      // 匿名导入,会调用对应包中的init函数，但不能使用包提供的接口函数
	mylib "gostudy/5-init/lib2"  // 别名导入 
)

func main() {
	mylib.Bar()
}
```

```go
package main

import (
	. "gostudy/5-init/lib2" // 使用 . 可以直接使用这个包里提供的接口函数（不建议使用）
)

func main() {
	Bar()
}
```

### 6 指针

略过 参考C语言中的指针

### 7 defer关键字

函数执行完后执行defer

```go
package main

import "fmt"

func main() {
   defer fmt.Println("go end")
   fmt.Println("fist go")
   fmt.Println("second go")
}
// ==>
// fist go
// second go
// go end
```

defer是以栈存储，运行时按照出栈顺序执行

```go
package main

import "fmt"

func main() {
   // defer执行顺序
   defer func1()
   defer func2()
   defer func3()
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
// ==>
// 3
// 2
// 1
```

defer在return执行后再执行

```go
package main

import "fmt"

func main() {
   func4()
}

func func4() int {
	defer fmt.Println("I'm defer")
	return func5()
}

func func5() int {
	fmt.Println("I'm return")
	return 0
}
// ==>
// I'm return
// I'm defer
```

### 8 数组与切片slice

#### 8.1 数组

```go
   // 固定长度的数组
   array1 := [10]int{1, 2, 3, 4}

   // 数组遍历方法一
   for i := 0; i < len(array1); i++ {
      fmt.Println(array1[i])
   }

   // 数组遍历方法二
   for index, value := range array1 {
      fmt.Println("index = ", index, "value = ", value)
   }

   // 动态数组 (常用)
   array2 := []int{11, 22, 33, 44}
```

- 固定数组再函数传参时严格匹配数组类型"[10]int"这种类似形式，且在传递时属于深拷贝

- 动态数组在函数传参是不匹配个数，且在传递时属于浅拷贝，传递的是数组的指针

#### 8.2 切片

```go
// 切片声明 
//方法一 这种方式不会给slice1分配空间  slice1==nil = true
var slice1 []int
slice1 = make([]int, 3) // 为slice1 分配3个空间，默认值为0
fmt.Printf("%v \n", slice1) // ==>[0 0 0]

// 方法二 
var slice1 []int = make([]int, 3)

// 方法三
slice1 := make([]int, 3)
```

##### 8.2.1切片容量追加

```go
// 切片容量追加
slice2 := make([]int, 3, 3)
fmt.Println(slice2)                       // ==> [0 0 0]
fmt.Println("len slice2 = ", len(slice2)) // ==> len slice2 =  3
fmt.Println("cap slice2 = ", cap(slice2)) // ==> cap slice2 =  3
slice2 = append(slice2, 2)
fmt.Println(slice2)                       // ==> [0 0 0 2]
fmt.Println("len slice2 = ", len(slice2)) // ==> len slice2 =  4
fmt.Println("cap slice2 = ", cap(slice2)) // ==> cap slice2 =  6
```

- 其中len是切片或数组的可见元素长度，cap是数组或切片的最大容量，每当apend操作后len超过cap后，go会自动执行 cap+=cap的操作,动态追加容量

- 在make时不指定cap值，以输入的len为默认cap值

##### 8.2.2 切片截取

```go
// 切片截取方法一 这种方式属于浅拷贝
slice3 := []int{1, 2, 3, 4}
s1 := slice3[1:3] // 左闭右开
s2 := slice3[:3]
fmt.Println(s1,s2) // ==> [2 3] [1 2 3]

// 切片截取方法二 copy(),这种方式需要提前给复制后的切片分配空间
s3 := make([]int, 4)
copy(s3, slice3)
fmt.Println(s3) // ==> [1 2 3 4] 
```

### 9 map

#### 9.1 map声明（参考[8.2切片声明]()）

- map的键值存储时为乱序，其原理采用哈希存储

```go
// 这里只展示一种 声明方式可参考切片声明
map1 := make(map[string]string, 5)
map1["a"] = "foo"
map1["b"] = "bar"
fmt.Println(map1) // ==> map[a:foo b:bar]
```

#### 9.1 map的使用方式

```go
// map 使用
// 增加
map1["c"] = "add"
// 修改
map1["a"] = "update"
// 删除
delete(map1, "b")
// 遍历
for key, value := range map1 {
   fmt.Println(key, value)
}
// ==>
// a update
// c add
```

- map 作为形参传递给函数时，传递的是map的指针

### 10 面向对象

#### 10.1 结构体

略过 参考C语言中的结构体

- 结构体作为形参传递给函数时，属于深拷贝

#### 10.2 对象封装与方法绑定

```go
func (receiver *Object) name()  {
	// do something
}
```



```go
package main

import "fmt"

// Dog 封装一个名为Dog的对象
type Dog struct {
	name string
}

// Eat 为Dog对象绑定方法
func (dog *Dog) Eat() {
	fmt.Println("dog eat shit!!!")
}

func (dog *Dog) getName() string {
	fmt.Println(dog.name)
	return dog.name
}

func (dog *Dog) setName(newname string) {
	// 要修改 dog对象时 需为接收器传入dog对象的指针 *Dog
	dog.name = newname
}
func main() {
	dog := Dog{name: "gg"}
	dog.getName()
	dog.setName("mm")
	dog.getName()
	dog.Eat()
}
```

#### 10.3 继承重写

```go
package main

import "fmt"

type Animal struct {
   name string
   legs int
}

func (animal Animal) Eat() {
   fmt.Println(animal.name, "can eat")
}

func (animal Animal) leg() {
   fmt.Println("animal have leg")
}

type dog struct {
   Animal
   food string
}

func (dog dog) Eat2() {
   fmt.Println(dog.name, "eat", dog.food)
}

func (dog dog) leg() {
   dog.Animal.leg()
   fmt.Println(dog.name, "有", dog.legs)
}
func main() {
   dog1 := Animal{
      name: "泰迪",
      legs: 4,
   }
   dog1.Eat() // ==> 泰迪 can eat

   dog2 := dog{
      Animal: Animal{
         name: "柴犬",
         legs: 4,
      },
      food: "屎",
   }
   dog2.Eat()  //父类的方法 ==> 柴犬 can eat
   dog2.Eat2() //子类的方法 ==> 柴犬 eat 屎

   var dog3 dog
   dog3.name = "哈士奇"
   dog3.legs = 4
   dog3.food = "沙发"

   dog3.Eat2() // 子类的方法 ==> 哈士奇 eat 沙发
   dog3.leg()  // 继承并重写父类的方法
}
```

#### 10.4 多态

- 需要有一个父类（提供接口）
- 需要有子类实现父类的全部接口方法
- 父类类型的变量（指针）指向（引用）子类的具体数据变量

```go
package main

import "fmt"

// AnimalIF interface本质是一个指针
type AnimalIF interface {
	aType() string
	aSleep()
}

type cat struct {
	catType string
}

// 子类实现父类接口方法
func (cat cat) aSleep() {
	fmt.Println("cat sleep")
}

func (cat cat) aType() string {
	return cat.catType
}

type dog struct {
	dogType string
}

func (dog dog) aSleep() {
	fmt.Println("dog sleep")
}

func (dog dog) aType() string {
	return dog.dogType
}

// 父类类型的变量（指针）指向（引用）子类的具体数据变量
func showType(animalIF AnimalIF) {
	fmt.Println(animalIF.aType())
}
func main() {
	var cat1 AnimalIF = &cat{"美短"}
	cat1.aSleep() // ==> cat sleep

	var dog1 AnimalIF = &dog{"土狗"}
	dog1.aSleep() // ==> dog sleep

	cat2 := cat{"银渐层"}
	showType(&cat2) // ==> 银渐层

	dog2 := dog{"沙皮狗"}
	showType(&dog2) // ==> 沙皮狗
}
```

#### 10.5 interface通用万能类型与interface断言

- 空接口
- int，string，struct ... 都实现了interface{}
- 可以用interface{}类型引用任意的数据类型

```go
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

// ==>
// arg type is main.book
// arg type is string
// arg type is string,value =  ttt
```





### **一些坑

1. 共有属性，私有属性

   想要其他包访问自己定义包的函数、对象等属性，需要将函数、对象等属性名称首字母大写。小写名称的函数、对象等属性只能在本包内使用

2. 自定义包

   - package 名必须与所在包名一致
   - 引入时路径基于系统环境变量的GOPATH路径