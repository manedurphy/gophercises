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
