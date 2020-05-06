package main

import "fmt"

func main() {
	// 定义一个变量
	var num1 int
	num1 = 30
	fmt.Printf("num1 的值为 %d \n", num1)
	// 定义的时候赋值
	var num2 = 40
	fmt.Printf("num2 的值为 %d \n", num2)
	// 简短声明
	str3 := "测试字符串"
	fmt.Printf("str3 的类型为 %T, 数值是 %s \n", str3, str3)

	// ---------------------------------------------------------------------------

	// 多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Printf("a= %d, b = %d, c= %d", a, b, c)

	var n1, n2, n3 = 100, 3.14, "你好"
	fmt.Printf("n1 = %d, n2 = %f, n3 = %s", n1, n2, n3)
}
