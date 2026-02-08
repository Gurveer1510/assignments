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

func ParseHTML2(r io.Reader) ([]AnchorTagData, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []AnchorTagData
	nodes := LinkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
		// fmt.Println(node)
	}

	return links, nil
}

func buildLink(n *html.Node) AnchorTagData {
	var ret AnchorTagData

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = getText2(n)

	return ret
}

func LinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, LinkNodes(c)...)
	}
	return ret
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
func getText2(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return "" 
	}
	var ret string
	for c := node.FirstChild ; c != nil; c= c.NextSibling {
		ret+= getText2(c) + " "
	}
	return ret
}
