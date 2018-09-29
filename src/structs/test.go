package main

import "fmt"
import "structs/computer"

func main(){	
	// import structs
	var spec computer.Spec
	spec.Maker = "apple"
	spec.price = 23423
	fmt.Println(spec)

}