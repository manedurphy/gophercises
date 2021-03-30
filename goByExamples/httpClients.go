package httpsClients

import (
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

	// scanner := bufio.NewScanner(resp.Body)

}
