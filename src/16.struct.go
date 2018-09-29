package main

import "fmt"

// type Employee struct {
// 	firstName string
// 	lastName string
// 	age int
// }

// Named Structure 
// type Employee struct {
// 	firstName, lastName string
// 	age, salary int
// }
// // Anonymous Structure
// var employee struct {
// 	firstName, lastName string
// 	age int
// }

// type Person struct {
// 	string 
// 	int
// }
// func main(){
// 	// create structure use field name
// 	emp1 := Employee{
// 		firstName: "wang",
// 		age: 23,
// 		salary: 34000,
// 		lastName: "gg",
// 	}
// 	// create structure without field name
// 	emp2 := Employee{"xue", "yuan", 25, 40000}
// 	fmt.Println(emp1)
// 	fmt.Println(emp2)

// 	emp3 := struct{
// 		firstName, lastName string
// 		age , salary int
// 	}{
// 		firstName:"yan",
// 		lastName: "xiao",
// 		age:38,
// 		salary:230000,
// 	}
// 	fmt.Println(emp3)

// 	// zero structure
// 	var emp4 Employee
// 	fmt.Println("zero structure is ", emp4)

// 	emp5 := Employee{
// 		firstName:"zhang",
// 		lastName:"long",
// 	}
// 	fmt.Println(emp5)
// 	// call structure body
// 	emp6 := Employee{"sun", "la", 23, 8999}
// 	fmt.Println("first name", emp6.firstName)
// 	fmt.Println("last name", emp6.lastName)
// 	fmt.Println("age", emp6.age)
// 	fmt.Println("salary", emp6.salary)
// 	// named from zero structure
// 	var emp7 Employee
// 	emp7.firstName = "he"
// 	emp7.lastName = "meng"
// 	fmt.Println("emp7 is ", emp7)
// 	// structure pointer
// 	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
// 	fmt.Println("first name ", (*emp8).firstName)
// 	fmt.Println("age",(*emp8).age)

// 	fmt.Println("first name",emp8.firstName)
// 	fmt.Println("age",emp8.age)

// 	// Anonymous Field
// 	p := Person{"game", 345}
// 	fmt.Println(p)
	
// 	var p1 Person
// 	p1.string = "gage"
// 	p1.int = 234
// 	fmt.Println(p1)
// }


// Nested Structure
// type Address struct {
// 	city, state string
// }

// type Person struct {
// 	name string
// 	age int
// 	address Address
// }

// func main(){
// 	var p Person
// 	p.name = "gan"
// 	p.age = 32
// 	p.address = Address{
// 		city: "beijinbg",
// 		state: "beijing",
// 	}
// 	fmt.Println(p)
// }

// Promoted Fields
// type Address struct {  
//     city, state string
// }
// type Person struct {  
//     name string
//     age  int
//     Address
// }

// func main(){
// 	var p Person
// 	p.name = "ggege"
// 	p.age = 23423
// 	p.Address = Address{
// 		city: "shijiaz",
// 		state: "bejing",
// 	}
// 	fmt.Println(p)
	
// 	// import structs
// 	var spec computer.Spec
// 	spec.Maker = "apple"
// 	spec.price = 23423
// 	fmt.Println(spec)

// }


// output struct body and code
// import "structs/computer"
// func main(){	
// 	// import structs
// 	var spec computer.Spec
// 	spec.Maker = "apple"
// 	spec.price = 23423
// 	fmt.Println(spec)
// }

// structs equality
type name struct {  
    firstName string
    lastName string
}

type image struct {
	data map[int]int
}

func main(){
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2{
		fmt.Println("name1 and name2 is equal")
	}else {
		fmt.Println("name1 is not equal name2")
	}
	name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
	}
	
	//如果结构体含有不可以比较的字段，则结构体变量也不可以比较
	image1 := image{data:map[int]int {
		2:123,
	}}
	image2 := image{data:map[int]int{
		0:155,
	}}

	if image1 == image2 {
        fmt.Println("image1 and image2 are equal")
    }
}