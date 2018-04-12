package main

import (
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// Link group info from a <a> tag
type Link struct {
	Href string
	Text string
}

func main() {

	var file string

	log.Println("######################")
	log.Println("[INF] Init application")
	log.Println("######################")

	file = "ex2.html"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("[ERR] Could not open file " + file)
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal("[ERR] Could not parse file " + file)
	}

	l := make([]Link, 0)
	log.Printf("%+v\n", findLinks(doc, &l))

}

func findLinkText(n *html.Node) string {
	ret := ""

	s := strings.TrimFunc(n.Data, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsPunct(r)
	})
	if n.Type == html.TextNode && s != "\n" {
		return s
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = ret + findLinkText(c)
	}

	return ret
}

func findLinks(n *html.Node, links *[]Link) []Link {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				linkText := findLinkText(n)
				l := Link{Text: linkText, Href: a.Val}
				*links = append(*links, l)
				return *links
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		*links = findLinks(c, links)
	}

	return *links
}
