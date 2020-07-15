package main

import (
	"fmt"
	"io"
	"os"
)

// create a program that reads a file on your harddrive and print the contents of the file to the terminal

// the file to open should be provided as an argument to the program when it is executed at the terminal.  For example, 'go run main.go myfile.txt should open up the myfile.txt file

// to read in the arguments provided to a program, you can reference the variable os.Args which is a slice of type string

// to open a file, check out the documentation for the Open function in the os package

// what interfaces does the file type implement

// if the file type implements the reader interfacew you might be able to reuse that io.copy function
type logWriter struct{}

var fileName string

func main() {
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	w := logWriter{}
	io.Copy(w, file)
}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
