package main

import (
	"fmt"
	"strings"

	links "github.com/manedurphy/gophercises/exercise4/link"
)

var html = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/first-page">A link to a <span>page</span></a>
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

	fmt.Printf("%+v\n", links)
}
