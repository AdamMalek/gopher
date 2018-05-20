package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gopher/zad4"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	domainFlag := flag.String("domain", "http://localhost:8080", "domain to scan")
	outputFilename := flag.String("output", "sitemap.xml", "sitemap xml filename")
	flag.Parse()

	parts := strings.Split(*domainFlag, "://")
	protocol := "http"
	var domain string
	if len(parts) > 1 {
		protocol = parts[0]
		domain = strings.Split(parts[1], "/")[0]
	} else {
		domain = strings.Split(*domainFlag, "/")[0]
	}
	domain = strings.TrimRight(domain, "/")

	baseURL := protocol + "://" + domain

	var visited []string
	var toVisit []string
	toVisit = append(toVisit, getRelativePath(baseURL))

	for len(toVisit) > 0 {
		curr := toVisit[0]
		curr = getRelativePath(curr)
		toVisit = toVisit[1:]
		visited = append(visited, curr)
		curr = baseURL + "/" + curr

		resp, err := http.Get(curr)
		if err != nil {
			fmt.Println("error requesting " + curr + ": " + err.Error())
		} else {
			links, err := linkparser.GetLinks(resp.Body)
			if err != nil {
				fmt.Println("error parsing " + curr)
			} else {
				for _, url := range links {
					if !isInDomain(url.URL, domain) {
						continue
					}
					urlRelative := getRelativePath(url.URL)
					newLink := true
					for _, v := range visited {
						if urlRelative == v {
							newLink = false
							break
						}
					}
					for _, v := range toVisit {
						if urlRelative == v {
							newLink = false
							break
						}
					}
					if newLink {
						fmt.Println("added: " + urlRelative)
						toVisit = append(toVisit, urlRelative)
					}
				}
			}
		}
	}

	err := export(*outputFilename, visited)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getRelativePath(url string) string {
	if strings.HasPrefix(url, "/") {
		return url
	}
	parts := strings.Split(url, "://")
	if len(parts) > 1 {
		parts = strings.Split(parts[1], "/")
	}
	if len(parts) > 1 {
		return "/" + strings.Join(parts[1:], "/")
	}
	return "/"
}

func isInDomain(url, domain string) bool {
	if strings.HasPrefix(url, "/") {
		return true
	}
	parts := strings.Split(url, "://")
	if len(parts) > 1 {
		parts = strings.Split(parts[1], "/")
	}
	if len(parts) > 1 {
		return parts[0] == domain
	}
	return false
}

func export(filename string, links []string) error {
	urlset := urlset{Namespace: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	for _, l := range links {
		urlset.Locs = append(urlset.Locs, xmlNode{URL: l})
	}

	xb, err := xml.Marshal(&urlset)
	if err != nil {
		return err
	}

	xmlBytes := append([]byte(xml.Header), xb...)

	err = ioutil.WriteFile(filename, xmlBytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

type urlset struct {
	Namespace string    `xml:"xmlns,attr"`
	Locs      []xmlNode `xml:"url"`
}

type xmlNode struct {
	URL string `xml:"loc"`
}
