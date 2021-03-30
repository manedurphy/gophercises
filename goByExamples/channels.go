package goByExamples

import "fmt"

func Channel() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func BufferedChannel() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	close(messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}
