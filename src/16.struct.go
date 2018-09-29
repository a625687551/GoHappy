package main

import "fmt"

// type Employee struct {
// 	firstName string
// 	lastName string
// 	age int
// }

// Named Structure 
type Employee struct {
	firstName, lastName string
	age, salary int
}
// Anonymous Structure
var employee struct {
	firstName, lastName string
	age int
}

type Person struct {
	string 
	int
}
func main(){
	// create structure use field name
	emp1 := Employee{
		firstName: "wang",
		age: 23,
		salary: 34000,
		lastName: "gg",
	}
	// create structure without field name
	emp2 := Employee{"xue", "yuan", 25, 40000}
	fmt.Println(emp1)
	fmt.Println(emp2)

	emp3 := struct{
		firstName, lastName string
		age , salary int
	}{
		firstName:"yan",
		lastName: "xiao",
		age:38,
		salary:230000,
	}
	fmt.Println(emp3)

	// zero structure
	var emp4 Employee
	fmt.Println("zero structure is ", emp4)

	emp5 := Employee{
		firstName:"zhang",
		lastName:"long",
	}
	fmt.Println(emp5)
	// call structure body
	emp6 := Employee{"sun", "la", 23, 8999}
	fmt.Println("first name", emp6.firstName)
	fmt.Println("last name", emp6.lastName)
	fmt.Println("age", emp6.age)
	fmt.Println("salary", emp6.salary)
	// named from zero structure
	var emp7 Employee
	emp7.firstName = "he"
	emp7.lastName = "meng"
	fmt.Println("emp7 is ", emp7)
	// structure pointer
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("first name ", (*emp8).firstName)
	fmt.Println("age",(*emp8).age)

	fmt.Println("first name",emp8.firstName)
	fmt.Println("age",emp8.age)
}
