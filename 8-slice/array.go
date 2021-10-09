package main

import "fmt"

func main() {
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

	// 动态数组切片 (常用)
	//array2 := []int{11, 22, 33, 44}

	// 切片声明
	//方法一 这种方式不会给slice1分配空间 slice1==nil = true
	var slice1 []int
	slice1 = make([]int, 3)     // 为slice1 分配3个空间，默认值为0
	fmt.Printf("%v \n", slice1) // ==>[0 0 0]

	// 方法二
	//var slice1 []int = make([]int, 3)

	// 方法三
	//slice1 := make([]int, 3)

	// 切片容量追加
	slice2 := make([]int, 3, 3)
	fmt.Println(slice2)                       // ==> [0 0 0]
	fmt.Println("len slice2 = ", len(slice2)) // ==> len slice2 =  3
	fmt.Println("cap slice2 = ", cap(slice2)) // ==> cap slice2 =  3
	slice2 = append(slice2, 2)
	fmt.Println(slice2)                       // ==> [0 0 0 2]
	fmt.Println("len slice2 = ", len(slice2)) // ==> len slice2 =  4
	fmt.Println("cap slice2 = ", cap(slice2)) // ==> cap slice2 =  6

	// 切片截取方法一 这种方式属于浅拷贝
	slice3 := []int{1, 2, 3, 4}
	s1 := slice3[1:3] // 左闭右开
	s2 := slice3[:3]
	fmt.Println(s1, s2) // ==> [2 3] [1 2 3]

	// 切片截取方法二 copy(),这种方式需要提前给复制后的切片分配空间
	s3 := make([]int, 4)
	copy(s3, slice3)
	fmt.Println(s3) // ==> [1 2 3 4]
}
