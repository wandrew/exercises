package main

import "fmt"

func main() {
	fmt.Println("This is a place to play with functions")
	name := "Bill"

	fmt.Println(&name)
	//fmt.Println(*name)
	fmt.Println(*&name)
}
