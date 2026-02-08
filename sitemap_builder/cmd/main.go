package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	sitemapbuilder "sitemap_builder"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value  string  `xml:"loc"`
}

type urlSet struct {
	Urls []loc	`xml:"url"`
	Xmlns	string	`xml:"xmlns,attr"`
}

func main() {
	requestedURL := flag.String("url", "https://www.sitemaps.org", "The URL for you which you want to generate the site builder")
	maxDepth := flag.Int("depth", 10, "The max depth you wanna go to find links.")
	flag.Parse()

	// pages := get(*requestedURL)
	pages := bfs(*requestedURL, *maxDepth)
	toXml := urlSet{
		Xmlns: xmlns,
	}
	

	for _, page := range pages {
		// fmt.Println(page)
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "   ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), WithPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := sitemapbuilder.ParseHTML2(r)
	var ret []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
			// default:
			// 	ret = append(ret, l.Href)
		}
	}

	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func bfs(urlStr string, maxDepth int) []string { 
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

func WithPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
