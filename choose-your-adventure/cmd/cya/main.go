package main

import (
	chooseyouradventure "choose_your_adventure"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start CYA web app on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Println("Using the story in", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := chooseyouradventure.JSONStory(f)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", story)

	tpl := template.Must(template.New("").Parse(storyTmpl))
	h := chooseyouradventure.NewHandler(
		story, 
		chooseyouradventure.WithTemplate(tpl), 
		chooseyouradventure.WithPathFunc(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	// h := chooseyouradventure.NewHandler(story, chooseyouradventure.WithTemplate(tpl))
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose your own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>
`
