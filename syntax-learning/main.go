package main

import (
	"fmt"
	"time"
)

var ( //这种因式分解的写法用来声明全局变量
	ab int
	ba bool
)

func main() {
	fmt.Println("Hello, World!")
	var number = 12
	var enddate = "2023-08-08"
	var url = "%d个人在%s学习"
	fmt.Printf(url, number, enddate)
	fmt.Println()
	fmt.Println("Google" + " Good")
	var targetstring = fmt.Sprintf(url, number, enddate)
	fmt.Println(targetstring)

	//数据类型
	// 布尔型bool
	// 数字类型 int float32 float64
	// 字符串类型 go的字符串是由单个字节链接起来的，字节用UTF-8编码标识Unicode文本
	// 派生类型 指针 数组 结构化类型 Channel类型 函数类型 切片类型 接口（interface）类型 Map类型
	// uintptr 无符号整型 用于存放指针

	//语言变量
	//未初始化即为0
	var a string = "learning"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d bool
	fmt.Println(d)
	var s string
	fmt.Println(s) //""空字符串
	//可以通过值自行判断
	var k = 1
	fmt.Println(k)
	// :=是声明语句 如果再加var反而会出bug 但只能再函数体中体现
	z := 1
	fmt.Println(ab, ba, z)
	// 并行赋值
	aa, bb, cc := 5, 8, "mirage"
	fmt.Println(aa, bb, cc)
	//空白标识符 _ 是一个只写变量 用于抛弃值
	_, b = 5, 7 // 5这个值被抛弃

	const lens = 10
	const wid = 5
	println(lens, wid)

	// iota 特殊变量 可以被认为是一个可以被编译器修改的常量
	// 在const关键字出现时被重置为一  每次用就加1
	const (
		ar = iota
		br = iota
		cr = iota
	)
	println(ar, br, cr)
	const (
		arr = 1 << iota
		brr = 3 << iota
		kr
		lr
	) //后面默认为3 再左移 iota位
	fmt.Println("arr=", arr)
	fmt.Println("brr=", brr)
	fmt.Println("kr=", kr)
	fmt.Println("lr=", lr)

	//运算符差不多
	//没有三目运算符

	// if
	// if+else
	// switch
	// select

	aa = 50
	if aa > 20 {
		fmt.Printf("aa的值大于20哟\n")
		fmt.Println("aa=", aa)
	}

	switch {
	case aa == 20:
		fmt.Println("aa=20")
	case aa > 20:
		fmt.Println("aa>20")
	case aa < 20:
		fmt.Println("aa<20")
	default:
		fmt.Println("这aa有问题")
	}
	// type switch
	// 可以用于判断interface变量中实际储存的变量类型
	var x interface{}
	switch i := x.(type) {
	case nil: //nil表示无值
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	fmt.Println()
	// fallthrough 加上后就会强制执行后面的语句，不判断了
	// 不然Go内部的 switch 自带着break
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough //也不管对不对直接执行2了
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	//select语句
	//select是Go中的一个控制结构 类似于switch语句
	//只能用于通道操作 每个case必须是一个通道操作 要么是发送要么是接收
	//监听所有指定通道上的操作 一旦其中一个准备好就执行相应代码块
	//多个都准备会随机执行一个通道 没有准备好的就执行default

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 循环语句只有for 和嵌套的for
	// break continue goto
	// 无限循环for true
	// for 三种用法
	// for init condition post
	// init赋值表达式 赋初值
	// condition关系表达式 循环控制
	// post赋值表达式 给控制变量或减量
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	su := 1
	for su <= 10 {
		su += su
	}
	fmt.Println(su)
	strings := []string{"yinsu", "mirage"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	numbers := [6]int{1, 2, 3, 5} //剩下都是0
	for i, x := range numbers {
		fmt.Printf("第%d位x的值=%d\n", i, x)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is %d -value is %f\n", key, value)
	}
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}
	for _, value := range map1 {
		fmt.Printf("value is %f\n", value)
	}
	var j int
	for i := 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}

	// select中使用break
	// 由于select是非阻塞的 会等待所有的通讯操作都准备就绪
	// 用 return 来提前结束执行的情况

	// select 我真他妈的不懂啊！！！！干
	// 下次再补着学吧

	// break 和 continue 中的  标记 re a
	// re会眺外层的 break直接结束 continue就在外层
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2 = %d\n", i2)
			break re
		}
	}
a:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2 = %d\n", i2)
			continue a
		}
	}

}

//	func funcname ( [parameter list] )[return_types]{
//		函数体
//	}
func max(n1, n2 int) int {
	var result int
	if n1 > n2 {
		result = n1
	} else {
		result = n2
	}
	return result
}

// 函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
