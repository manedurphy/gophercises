package goByExamples

import (
	"fmt"
	"net/http"
)

var urls []string = []string{"https://google.com", "https://twitter.com"}

func workerWG(url string, respCh chan<- *http.Response) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	respCh <- resp

}

func FetchStatus(w http.ResponseWriter, r *http.Request) {
	// var wg sync.WaitGroup
	respCh := make(chan *http.Response, 1)

	for _, url := range urls {
		// wg.Add(1)
		go workerWG(url, respCh)
		// wg.Done()
	}
	// wg.Wait()

	resp1 := <-respCh
	// resp2 := <-respCh

	// fmt.Fprintf(w, "%+v\n", resp1.Status)
	// fmt.Fprintf(w, "%+v\n", resp2.Status)
	fmt.Printf("%+v\n", resp1.Status)
	// fmt.Printf("%+v\n", resp2.Status)
}
