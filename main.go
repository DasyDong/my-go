package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0], a[1])   // 打印数组的前2个元素
	fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

}
