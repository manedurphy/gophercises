package exercise3

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
}

func GetDecodedJson(file *os.File) (Story, error) {
	d := json.NewDecoder(file)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story) handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(htmlTemplate))

	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

var htmlTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Choose Your Own Adventure</title>
</head>
<body>
	<h1>{{.Title}}</h1>
	{{range .Paragraphs}}
	<p>{{.}}</p>
	{{end}}
	<ul>
		{{range .Options}}
		{{.}}
		<li>
			<a href="/{{.Chapter}}">{{.Text}}</a>
		</li>
		{{end}}
	</ul>
</body>
</html>
`
