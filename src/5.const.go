package main


// 定义
// import (
// 	"fmt"
// 	"math"
// )

// func main(){
// 	const a = 89
// 	// a = 33 // forbid to recover
// 	fmt.Println("begin our processing")
// 	var b = math.Sqrt(4)
// 	// const c = math.Sqrt(4) // 不允许这样,常量需要编译时候就需要确定，而函数在编译后才知道结果

// }

// string const
// import "fmt"

// func main(){
// 	var name = "sma"
// 	fmt.Println("type %T value %v", name, name)  // 这句话就不起作用
// 	fmt.Printf("type %T value %v", name, name)

// 	var defaultName = "jame"
// 	type myString string
// 	var cutomeName myString = "jame"
// 	// cutomeName = defaultName  // 这个不允许
// }


// 数字表达式

import "fmt"
func main(){
	var a  = 4.4/3
	fmt.Println("value of %v", a)
}