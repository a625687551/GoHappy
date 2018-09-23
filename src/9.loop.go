package main

/* for is only loop in golang
for initialisation; condition; post {
}
*/

import "fmt"

func main(){
	for i:= 1;i<=10;i++{
		fmt.Print(i)
	}
	fmt.Println("second begin")
	for i := 1;i<=10;i++{
		if i >5{
			break
		}
		fmt.Println(i)
	}
	fmt.Print("line after for loop, begin to continue")

	for i := 1;i<=10;i++{
		if i%2 == 0{
			continue
		}
		fmt.Print("\n", i)
	}

	j := 1
	for ;j<=10;{
		fmt.Print("j ", j)
		j += 2
	}

}