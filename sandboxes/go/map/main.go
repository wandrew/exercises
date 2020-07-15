package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"white": "ffffff",
		"black": "000000",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for clr, hex := range c {
		fmt.Printf("%v's hex code is %v\n", clr, hex)
	}
}
