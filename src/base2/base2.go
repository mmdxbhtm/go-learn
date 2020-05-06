package main

import "fmt"

var a = 1000

func main() {
	var num = 100
	fmt.Printf("num 的数值是 %d， num的地址是 %p \n", num, &num)
	num = 200
	fmt.Printf("num 的数值是 %d， num的地址是 %p \n", num, &num)

	fmt.Println(a)

	var a1 int
	fmt.Println(a1)

	var f1 float32
	fmt.Println(f1)

	var str1 string
	fmt.Println(str1)

	var arr []int
	fmt.Println(arr)

}
