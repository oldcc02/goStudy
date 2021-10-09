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
