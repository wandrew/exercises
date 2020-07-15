package main

import (
	"fmt"
	"net/http"
	"time"
)

type (
	worker struct {
		c chan string
	}
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	// Start our routines
	// how are addresses for s assigned over range?  If you change the function literal to s *string it only prints out
	//  amazon.
	for _, s := range sites {
		go func(url string) { // <- this is our parameter definition (function signature)
			for {
				c <- getStatus(url)
				time.Sleep(10 * time.Second)
			}
		}(s) // <--- this calls the function literal so 's' is the value we are passing
	}

	// monitor our channels
	for l := range c {
		fmt.Println(l)
	}
}

func getStatus(url string) string {
	_, err := http.Get(url)
	if err == nil {
		return fmt.Sprintf("%v appears to be up", url)
	}
	return fmt.Sprintf("%v might be down with error: %v", url, err)

}
