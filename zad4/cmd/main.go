package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopher/zad4"
)

func main() {

	file := flag.String("file", "ex1.html", "file to parse")
	flag.Parse()

	path, _ := (filepath.Abs("../examples/" + *file))
	r, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: " + path)
	} else {
		links := linkparser.GetLinks(r)
		fmt.Printf("%+v\n", links)
	}
}
