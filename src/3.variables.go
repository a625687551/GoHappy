package main

import "fmt"

func main(){
	// var age int // 声明变量,如果没有赋值，则会自动赋值一个初始值
	// fmt.Println("my age is", age)
	// age = 28
	// fmt.Println("my age is", age)
	// age = 54
	// fmt.Println("my age is ", age)
	// var day int = 34
	// fmt.Println("my day is ", day)	
	// var time = 12 //有自动的类型推断
	// fmt.Println("time is ", time)
	// // var width, height int //忽略多变量初始化
	// var width, height = 100, 200 //多变量声明
	// fmt.Println("width si ", width, "height is ", height)

	var (
		name = "maomaochong"
		age = 28
		height int 
	)
	fmt.Println("my name is ", name, "age is ", age, "and heigt is ", height)
}