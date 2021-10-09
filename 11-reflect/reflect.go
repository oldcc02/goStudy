package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	name     string
	readTime int
}

func (book *Book) readBook(time int) {
	fmt.Println("read a book")
	book.readTime += time
}

func myReflect(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func reflectFieldAndMethod(arg interface{}) {
	// 获取arg的type
	inputType := reflect.TypeOf(arg)
	fmt.Println("input arg type is", inputType)

	// 获取arg的value
	inputValue := reflect.ValueOf(arg)
	fmt.Println("input arg value is", inputValue)

	// 通过inputType.Field(i)获取arg的字段，字段的数据类型
	// 通过inputValue.Field(i).Interface()获取arg字段的值
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s %v = %v \n", field.Name, field.Type, value)
	}

	//通过inputType.Method(i)获取arg的方法调用
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s %T \n", method.Name, method.Type)
	}
}

func main() {
	var str string = "tttt"
	myReflect(str)

	var book Book = Book{
		name:     "哈利波特",
		readTime: 10,
	}
	//myReflect(book)
	reflectFieldAndMethod(book)
}
