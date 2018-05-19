package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"gopher/zad4"
)

func main() {

	file := flag.String("file", "", "file to parse")
	url := flag.String("url", "", "url to get links from, has highrer priority over --file. Format: http(s)://xxx.xxx")
	flag.Parse()

	var r io.Reader
	var err error
	r = nil
	if *url != "" {
		resp, err := http.Get(*url)
		if err != nil {
			fmt.Println("Error getting html from URL: " + *url)
		} else {
			r = resp.Body
		}
	} else if *file != "" {
		r, err = os.Open("../examples/" + *file)
		if err != nil {
			p, _ := filepath.Abs("../examples/" + *file)
			fmt.Println("Error opening file: " + p)
		}
	} else {
		fmt.Println("No source selected")
		return
	}

	links, err := linkparser.GetLinks(r)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("%+v\n", links)
	}
}
