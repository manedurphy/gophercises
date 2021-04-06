package main

import (
	"fmt"
	"time"

	"github.com/manedurphy/gophercises/goByExamples"
)

var urls []string = []string{"https://google.com", "https://twitter.com"}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go goByExamples.MakeRequest(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
