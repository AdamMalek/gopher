package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// GetLinks gets all links found
func GetLinks(r io.Reader) ([]Link, error) {

	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	addNode := func(l Link) {
		links = append(links, l)
	}

	parse(node, addNode)

	return links, nil
}

func parse(node *html.Node, addLink func(Link)) {
	if node.Type == html.ElementNode && node.Data == "a" {
		addLink(parseLink(node))
	} else {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			parse(c, addLink)
		}
	}
}

func parseLink(node *html.Node) Link {
	url := ""
	for _, a := range node.Attr {
		if a.Key == "href" {
			url = strings.Split(a.Val, "?")[0]
			url = strings.Split(url, "#")[0]
		}
	}

	return Link{URL: url, Text: getText(node)}
}

func getText(node *html.Node) string {
	text := ""
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		} else if c.Type == html.ElementNode {
			text += getText(c)
		}
	}
	return strings.TrimSpace(text)
}
