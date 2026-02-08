package htmllinkparser

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type AnchorTagData struct {
	Href string
	Text string
}


func ParseHTML(r io.Reader) []AnchorTagData {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	var tags []AnchorTagData
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					atd := AnchorTagData{a.Val, getText(n)}
					tags = append(tags, atd)
				}
			}
		}
	}
	return tags
}

func getText(node *html.Node) string {
	var text string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			text += strings.TrimSpace(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)
	return text
}
