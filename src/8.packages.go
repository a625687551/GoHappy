package main

/*
if condition {
	pass
}

if condition{

}else if condition{

}else {

}
*/

import "fmt"


func main(){
	num := 10
	if num%2 == 0{
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	if num1 := 23; num1 % 2 == 0{
		fmt.Println("is even")
	} else {
		fmt.Println("number is odd")
	}

	if num <= 50{
		fmt.Println("number is less than  or equal to 50")
	} else if num >= 51 && num <= 100{
		fmt.Println("nuber is between 51 and 100")
	} else {
		fmt.Println("number is more than 100")
	}
}

