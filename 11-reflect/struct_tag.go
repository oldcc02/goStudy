package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `info:"姓名"`
	Sex  string `info:"性别"`
}

func getTag(i interface{}) {
	inputType := reflect.TypeOf(i)
	tagKey := "info"
	for i := 0; i < inputType.NumField(); i++ {
		tagValue := inputType.Field(i).Tag.Get(tagKey)
		fmt.Println(tagKey, tagValue)
	}
}

func main() {
	var xiaoHong Student
	getTag(xiaoHong)
}

//info 姓名
//info 性别
