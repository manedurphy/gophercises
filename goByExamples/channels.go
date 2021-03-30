package goByExamples

import (
	"fmt"
	"time"
)

func Channel() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()
	go func() {
		messages <- "pong"
	}()

	msg1 := <-messages
	msg2 := <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)
}

func BufferedChannel() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// messages <- "third" // this will cause infinite wait
	close(messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done!")

	done <- true
}

func Synchronization() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func PingPong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
