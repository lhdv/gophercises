package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/lhdv/gophercises/02_urlshortener/urlshort"
)

func main() {

	var linksFile string
	var customHandler http.HandlerFunc

	flag.StringVar(&linksFile, "i", "links.yaml", "links file(yaml or json) with url and links")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	content := readFile(linksFile)

	customHandler, err := handlerChooser(linksFile, content, &mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", customHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readFile(name string) string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("Can't read yaml file")
	}

	return string(b)
}

func handlerChooser(fileName, fileContent string, fallback *http.HandlerFunc) (http.HandlerFunc, error) {
	if filepath.Ext(fileName) == ".json" {
		return urlshort.JSONHandler([]byte(fileContent), *fallback)
	}

	return urlshort.YAMLHandler([]byte(fileContent), *fallback)
}
