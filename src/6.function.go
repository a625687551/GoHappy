package main

import (
	"fmt"
)

/*
func functionname(parametername type) returntype {
	//function body
}
*/

//简单尝试
// func colculateBill(price int, no int) int {
// func colculateBill(price, no int) int {
// 	totalPrice := price * no
// 	return totalPrice
// }

// func main(){
// 	price , no := 3, 4
// 	totalPrice := colculateBill(price, no)
// 	fmt.Println("total price is ",totalPrice)
// }


// multi return
// func rectProps(length, width float64)(float64, float64){
// 	area := length*width
// 	perimeter := (length + width)*2
// 	return area, perimeter
// }

// func main(){
// 	area, per := rectProps(3.34, 8.3)
// 	fmt.Printf("area %f, perimeter %f", area, per)
// }


// rename return name

// func rectProbs(length, width float64)(area, perimeter float64){
// 	area = length*width
// 	perimeter = 2*(length + width)
// 	return
// }
// func main(){
// 	area, per := rectProbs(3.34, 8.3)
// 	fmt.Printf("area %f, perimeter %f", area, per)
// }


// empty sybol
func rectProbs(length, width  float64)(area, perimeter float64){
	area = length*width
	perimeter = 2*(length + width)
	return
}

func main(){
	area, _ := rectProbs(4.3, 8.6)
	fmt.Printf("area is %v", area)
}