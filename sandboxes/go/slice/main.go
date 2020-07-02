package main

import "fmt"

func main() {
	// Slice -> {pointer-to-head,capacity,lenth} -> array{hi, there, how, are you}
	//             |_______________________________________^

	mySlice := []string{"Hi", "There", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

// gotcha because no pointer :-D
// the slice is passed by value, but because a slice is just ptr,cap,len,
// it is still modifying the same array that both copies of the slice point to ;-)
//
// this applies to maps, channels, pointers, and functions (this will be fun)
func updateSlice(s []string) {
	s[0] = "Bye"
}
