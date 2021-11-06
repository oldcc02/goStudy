package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name     string    `json:"name"`
	Sex      string    `json:"sex"`
	Age      int       `json:"age"`
	Students []student `json:"students"`
}

type student struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func main() {
	xiaoHong := student{
		Name: "小红",
		Sex:  "女",
	}
	xiaoMing := student{
		Name: "小明",
		Sex:  "男",
	}
	laoZhang := Teacher{
		Name:     "老张",
		Sex:      "男",
		Age:      20,
		Students: []student{xiaoHong, xiaoMing},
	}
	// 将结构体转换为json
	jsonResult, err := json.Marshal(laoZhang)
	if err != nil {
		fmt.Printf("err %s", err)
		return
	}
	fmt.Printf("%s \n", jsonResult)

	// 将json转换为结构体
	teacherStruct := Teacher{}
	err = json.Unmarshal([]byte("{\"name\":\"老张\",\"sex\":\"男\",\"age\":20}"), &teacherStruct)
	if err != nil {
		fmt.Printf("err %s", err)
	}
	fmt.Printf("%v \n", teacherStruct)
}

// ==>
//{"name":"老张","sex":"男","age":20,"students":[{"name":"小红","sex":"女"},{"name":"小明","sex":"男"}]}
//{老张 男 20 []}
