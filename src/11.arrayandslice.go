package main

import "fmt"

// func main(){
// 	var a [3]int
// 	fmt.Println(a)
// 	a[0] = 12
// 	a[1] = 23
// 	a[2] = 33
// 	fmt.Println(a)
// 	b := [3]int{23,34,78}
// 	fmt.Println(b)
// 	c := [3]int{13}
// 	fmt.Println(c)
// 	d := [...]int{3,4,9,8,9}
// 	fmt.Println(d)
// 	e := [...]string{"USA", "China", "India", "Germany", "France"}
// 	f := e
// 	f[0] = "wangjifeng"
// 	fmt.Println(e)
// 	fmt.Println(f)
// 	fmt.Println(len(f))
// 	//iter array
// 	for i :=0;i<len(f);i++{
// 		fmt.Printf("%d th element of is %v \n",i, f[i])
// 	}
// 	for i, v := range(f){
// 		fmt.Printf("%d th element of is %v \n",i, v)
// 	}
// 	//多维数组pass

// }


// slice have not own data is just data mirror
// func main(){
// 	a := [5]int{76, 77, 78, 79, 80}
// 	var b []int = a[1:4]
// 	fmt.Println(b)
// 	c := []int{5, 6, 7}
// 	fmt.Println(c)
// 	origin := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
// 	middle := origin[2:5]
// 	fmt.Println(origin)
// 	for i := range(middle){
// 		middle[i]++
// 	}
// 	fmt.Println(origin)
// 	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
// 	fruitslice := fruitarray[1:3]
// 	fmt.Println(fruitslice)
// 	fmt.Printf("length of slice %d capacity %d \n", len(fruitslice), cap(fruitslice))
// 	fruitslice = fruitslice[:cap(fruitslice)]
// 	fmt.Println(fruitslice)
// 	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
// }

// creat slice
func main(){
	i := make([]int, 5, 5)
	fmt.Println(i)
	// 当存储空间不足时候，每次slice的容量都会增加一倍，直到下次容量不足
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyata")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "good")
	cars = append(cars, "good1")
	cars = append(cars, "good2")
	cars = append(cars, "good3")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	var names []string
	if names == nil{
		fmt.Println("slice is nil going to append")
		names = append(names, "wang","xuey", "xiaohe")
		fmt.Println("names", names, len(names), cap(names))
	}
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println(food)
}
