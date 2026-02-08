package main

import (
	"fmt"
	htmllinkparser "html_link_parser"
	"os"
)

func main() {
	// fmt.Println("hello")
	file, err := os.Open("index2.html")
	if err != nil {
		panic(err)
	}

	tags, _ := htmllinkparser.ParseHTML2(file)

	for _, v := range tags {
		fmt.Printf("%v -------------- %v\n", v.Href, v.Text)
	}

}
