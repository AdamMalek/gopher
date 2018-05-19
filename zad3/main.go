package main

import (
	"flag"
	"fmt"
	"net/http"

	"gopher/zad3/handler"
	"gopher/zad3/parser"
)

func main() {
	file := flag.String("file", "story.json", "JSON file containing story")
	port := flag.String("port", "8090", "port")
	flag.Parse()

	storyProvider := parser.CreateProvider(*file)
	handler := handler.GetHandler(storyProvider, defaultMux())

	fmt.Println("Starting the server on :" + *port)

	http.ListenAndServe(":"+*port, handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", 200)
	})
	return mux
}
