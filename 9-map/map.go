package main

import "fmt"

func main() {
	// map 声明
	map1 := make(map[string]string, 5)
	map1["a"] = "foo"
	map1["b"] = "bar"
	fmt.Println(map1) // ==> map[a:foo b:bar c:pee]

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
}
