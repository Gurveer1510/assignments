package main

import (
	"fmt"
	htmllinkparser "html_link_parser"
	"os"
)

func main() {
	// fmt.Println("hello")
	file, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}

	tags := htmllinkparser.ParseHTML(file)

	for _, v := range tags {
		fmt.Printf("%v -------------- %v\n", v.Href, v.Text)
	}

}
