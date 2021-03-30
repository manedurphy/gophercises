package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manedurphy/gophercises/exercise3"
)

func main() {
	port := flag.Int("port", 3000, "port to start server")
	jsonFileName := flag.String("json", "gopher.json", "a json file with story arcs")
	flag.Parse()

	file, err := os.Open(*jsonFileName)
	if err != nil {
		panic(err)
	}

	story, err := exercise3.GetDecodedJson(file)
	if err != nil {
		panic(err)
	}

	h := exercise3.NewHandler(story)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
