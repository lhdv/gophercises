package main

import (
	"log"
	"os"
	"strings"

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

	file = "ex3.html"
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatal("[ERR] Could not open file " + file)
	}

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal("[ERR] Could not parse file " + file)
	}

	ll := make([]Link, 0)
	ll = findLinks(doc, &ll)
	for _, l := range ll {
		log.Printf("%+v\n", l)
	}

}

func findLinks(n *html.Node, links *[]Link) []Link {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				linkText := findLinkText(n)
				l := Link{Href: a.Val, Text: linkText}
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

func findLinkText(n *html.Node) string {
	ret := ""

	// 1 - Split a blank separated string into slice
	// 2 - Join this slice using a separator
	s := strings.Join(strings.Fields(n.Data), " ")

	if n.Type == html.TextNode && s != "\n" {
		return s
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = ret + findLinkText(c)
	}

	return ret
}
