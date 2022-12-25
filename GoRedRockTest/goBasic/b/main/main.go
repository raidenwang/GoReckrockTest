package main

import "fmt"

/*
错误原因：被 defer修饰的func（），将会在函数返回之前执行。最后一个defer前函数已返回，故不会执行最后一个func（）
*/
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
