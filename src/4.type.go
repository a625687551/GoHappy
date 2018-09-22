package main
/*
bool
数字类型
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
string
*/

// import "fmt"

// func main(){
// 	fmt.Println("begin")
// 	a ,b := true, false
// 	fmt.Println("a: ", a, "b: ", b)
// 	c := a && b
// 	fmt.Println("c :", c)
// 	// var d bool = a || b
// 	d := a || b
// 	fmt.Println("d: ", d)
// }


// 有符号整型
// import (
// 	"fmt"
// 	"unsafe"
// )

// func main(){
// 	var a int = 39
// 	b := 93
// 	fmt.Println("value of a is ", a, "valule of b is ", b)
// 	fmt.Println("type of a is %T", a, "size of a is %T", unsafe.Sizeof(a))
// 	fmt.Println("type of a is %T", b, "size of a is %T", unsafe.Sizeof(b))
// }

//无符号整型和浮点数
// import "fmt"

// func main(){
// 	a, b := 3.43, 8.33
// 	fmt.Println("type of a %T, b %T", a, b)
// 	sum := a + b
// 	diff := a - b
// 	fmt.Println("sum:", sum, "diff:", diff)

// 	no1, no2 := 56, 89
// 	fmt.Println("sum", no1+no2, "diff: ", no1-no2)
// }


//复数类型
// import "fmt"

// func main(){
// 	c1 := complex(3, 4)
// 	c2 := 2 + 9i
// 	cadd := c1 + c2
// 	fmt.Println("sum :", cadd)
// 	cmu := c1*c2
// 	fmt.Println("product :", cmu)
// }

// string类型
// import "fmt"

// func main(){
// 	first := "wang"
// 	last := "chong"
// 	name := first + " " + last
// 	fmt.Println("name is ", name)
// }

//类型转换: go里面的类型转换比较严格，没有自动提升类型和类型转换
import "fmt"

func main(){
	i := 3.33
	j := 232
	// sum := i + j // 不许int + float
	sum := int(i) + j
	fmt.Println("sum is ", sum)
	var k float64 = float64(j)
	fmt.Println("type of k is %T", k) //不知道为啥这里又错误
}