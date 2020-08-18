package main

import "fmt"

func main() {
	b := Book("On the Road")
	fmt.Println(b) // 打印或输出为字符串时会调用String()方法
}

type Book string

func (b Book) String() string {
	return string("《" + b + "》")
}
