package goByExamples

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Routine() {
	f("direct")

	go f("routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func Sample() {
	go func() {
		fmt.Println("Box 1")
		fmt.Println("Box 2")
		fmt.Println("Box 3")
		fmt.Println("Box 1")
		fmt.Println("Box 2")
		fmt.Println("Box 3")
		fmt.Println("Box 1")
		fmt.Println("Box 2")
		fmt.Println("Box 3")
	}()

	fmt.Println("Box 4")
	fmt.Println("Box 5")
	fmt.Println("Box 6")
	fmt.Println("Box 7")
	fmt.Println("Box 8")
	fmt.Println("Box 9")
	fmt.Println("Box 10")
	time.Sleep(time.Second) // without this the function running the main logic finishes before the go routine, and exits before it can log
}

func task(arg int) {
	fmt.Println("Box", arg)
}

func RoutineLoop() {
	var N int = 100

	for i := 0; i < N; i++ {
		task(i)
	}

}

func GoRoutineLoop() {
	var N int = 100

	acknowledgeChannel := make(chan bool, 100)

	for i := 0; i < N; i++ {
		go func(i int) {
			task(i)
			acknowledgeChannel <- true
		}(i)
	}

	for i := 0; i < N; i++ {
		<-acknowledgeChannel
	}
}

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
