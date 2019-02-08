// package sitemap
package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Link group info from a <a> tag
type Link struct {
	Href string
	Text string
}

func main() {

	var site string
	var depth int
	var sitemap map[string]string
	var level1 map[string]string

	flag.StringVar(&site, "url", "https://lhdv.net", "the base url to get sitemap")
	flag.IntVar(&depth, "depth", 0, "maximum depth you will go down to build sitemap")

	flag.Parse()

	// sitemap = fetchLinks(site)
	// for k := range sitemap {
	// 	level1 = fetchLinks(k)
	// 	log.Println(k, "->", level1)
	// }

}

func fetchLinks(site string) map[string]string {

	sitemap := make(map[string]string)

	domain, err := parseDomain(site)
	if err != nil {
		log.Fatal("[ERR][parseDomain] - ", err)
	}

	resp, err := http.Get(site)
	if err != nil {
		log.Fatal("[ERR][http.Get] - ", err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("[ERR][html.Parse] - ", err)
	}

	ll := make([]Link, 0)
	ll = findLinks(doc, &ll)
	for _, l := range ll {
		link, err := parseDomain(l.Href)
		if err != nil {
			log.Fatal("[ERR][parseDomain] - ", err)
		}

		if link == domain {
			sitemap[l.Href] = l.Text
		}

	}

	// log.Println(doc, depth)

	return sitemap
}

func parseDomain(site string) (string, error) {
	site = strings.TrimSpace(site)
	urlParsed, err := url.Parse(site)
	return urlParsed.Hostname(), err
}

//
// From previous exercise
//
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
