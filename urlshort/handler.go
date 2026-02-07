package main

import (
	"fmt"
	"net/http"

	// "github.com/goccy/go-yaml"
	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		redirect_url := r.URL.Path
		// fmt.Println(redirect_url)
		if url, ok := pathsToUrls[redirect_url]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type yamlStruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var m []yamlStruct

	err := yaml.Unmarshal(yml, &m)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path
		fmt.Println("here in 1 log", shortUrl)
		fmt.Println(m)
		for _, v := range m {
			if v.Path == shortUrl {
				fmt.Println(v.Path)
				fmt.Println("here")
				http.Redirect(w, r, v.Url, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
