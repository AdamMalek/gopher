package handler

import (
	"html/template"
	"net/http"
	"strings"

	"gopher/zad3/parser"
)

// GetHandler returns http handler
func GetHandler(parser parser.Provider, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		story, exists := parser(strings.TrimLeft(request.URL.Path, "/"))
		if exists {
			tmp, _ := template.ParseFiles("html/view.html")
			tmp.Execute(rw, story)
		} else {
			fallback.ServeHTTP(rw, request)
		}
	}
}
