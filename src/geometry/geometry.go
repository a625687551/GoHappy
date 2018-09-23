// geometry.go
package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

func init(){
	fmt.Println(" main package initalized")
	if len<0 {
		log.Fatal("len is less than zero")
	if wid < 0 {
		log.Fatal("wid is less than zero")
	}
	}
}
func main(){
	var len, wid float64 = 4, 3
	fmt.Println("Geometry  shape properties")
	fmt.Printf("area of rectangle %.2f \n", rectangle.Area(len, wid))
	fmt.Printf("diagonal of rectangle %.2f \n", rectangle.Diagonal(len, wid))
}