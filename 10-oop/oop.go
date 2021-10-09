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
