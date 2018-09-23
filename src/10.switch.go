package main

import "fmt"

// case wil not allow same case value
func main(){
	finger := 6
	switch finger{
	case 1:
		fmt.Print("thumb")
	case 2:
		fmt.Print("index")
	case 3:
		fmt.Print("middle")
	case 4:
		fmt.Print("ring")
	case 5:
		fmt.Print("pinky")
	default:
		fmt.Print("which figer you have, six?")
	}
	
	letter := "i"
	switch letter{
	case "a", "i", "d":
		fmt.Print("\nvowel")
	default:
		fmt.Print("\nnot a vowel")
	}

	num := 43
	switch{
	case num >=0 && num <= 50:
		fmt.Print("number is in [0,50]")
	case num >=51&& num <=100:
		fmt.Print("number is in [51,100]")
	default:
		fmt.Print("num maybe greater than 101")
	}

	switch{
	case num <50:
		fmt.Println("number is less than 50")
		fallthrough
	case num <100:
		fmt.Println("umber is less than 100")
		fallthrough
	case num <200:
		fmt.Println("number is less than 200")
		fallthrough
	default:
		fmt.Println("you win...")
	}
}