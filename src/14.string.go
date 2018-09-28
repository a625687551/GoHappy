package main

import "fmt"

func main(){
	name := "hello world"
	fmt.Println(name)
	printBytes(name)
	printLettle(name)
	fmt.Println("char print")
	printChar(name)
	// create string from bytes
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str1 := string(runeSlice)
	fmt.Println(str1)
	fmt.Println(len(str))

}

func printBytes(s string){
	for i, value := range(s){
		fmt.Printf(" key %v value %v ", i, value)
		fmt.Printf(" key %x value %x ", i, value)
		// fmt.Printf("value %c ", value)
	}
}


func printLettle(s string){
	for i := range s{
		fmt.Printf("value is %c", s[i])
	}
}

func printChar(s string){
	runes := []rune(s)
	for _, value := range runes{
		fmt.Printf(" %c ", value)
	}
}