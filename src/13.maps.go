package main

import (
	"fmt"
)


func main(){
	var personsalary map[string]int
	if personsalary == nil{
		fmt.Println("map is empty going to add one")
		personsalary = make(map[string]int)
		fmt.Println(personsalary)
	}
	personsalary["wang"] = 230000
	fmt.Println(personsalary)
	// one way
	job1 := make(map[string]int)
	job1["data"]= 2342
	fmt.Println(job1)
	// two
	job := map[string]int{
		"minig":23234234,
		"ai": 23000,
	}
	job["data"] = 13000
	fmt.Println(job)
	value, ok := job["ai"]
	if ok == true{
		fmt.Println("ai is in ", value)
	}else{
		fmt.Println("ai is not in")
	}
	//iter
	for key,value := range job{
		fmt.Println("the key value is ", key, value)
	}
	//delete element
	delete(job, "minig")
	fmt.Println(job)
	//length 
	fmt.Println("job length is ", len(job))
	//map 也是引用类型,map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
}
