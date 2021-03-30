package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/manedurphy/gophercises/exercise3"
)

func main() {
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

	fmt.Printf("%+v", story)
}

// type Story map[string]Chapter

// type Chapter struct {
// 	Title      string   `json:"title"`
// 	Paragraphs []string `json:"story"`
// 	Options    []Option `json:"options"`
// }

// type Option struct {
// 	Text    string `json:"text"`
// 	Chapter string `json:"arc"`
// }

// func getDecodedJson(file *os.File) (Story, error) {
// 	d := json.NewDecoder(file)

// 	var story Story
// 	if err := d.Decode(&story); err != nil {
// 		return nil, err
// 	}
// 	return story, nil
// }
