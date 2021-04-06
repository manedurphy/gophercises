package main

import (
	"fmt"
	"strings"

	"github.com/manedurphy/gophercises/exercise4/links"
)

var html = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(html)
	links, err := links.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}
