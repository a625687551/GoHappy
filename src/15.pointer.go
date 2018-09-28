package main

import "fmt"

// func main(){
// 	b := 344
// 	var a *int = &b
// 	fmt.Printf("type of a is %T ", a)
// 	fmt.Println("address of a is ", a)
// }

// zero pointer
// func main(){
// 	a := 346
// 	var b *int
// 	if b == nil {
// 		fmt.Println("b is nil")
// 		b = &a
// 		fmt.Println(b, "is a address")
// 		fmt.Println(*b, "is get value of a from pointer")
// 		*b++
// 		fmt.Println("value of a is change to ", a)
// 		change(b)
// 		fmt.Println("b after changed function is ", a)
// 	}
// }

// func change(val *int){
// 	*val = 55
// }


// 不要向函数传递数组的指针，而应该使用切片
// func main(){
// 	a := [...]int{23,34, 44}
// 	fmt.Println(a)
// 	modify(a[:])
// 	fmt.Println(a)
// }
// func modify(s []int){
// 	s[0] = 99
// }

// Go 并不支持其他语言（例如 C）中的指针运算
func main(){
	b := [...]int{23,3434,343,2}
	p := &b
	fmt.Println("value b and p is ", b, p)
	p++
}