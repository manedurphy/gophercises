package goByExamples

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpClients() {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", resp.Body)

	scanner := bufio.NewScanner(resp.Body)
	fmt.Println("Scanner:", scanner)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
