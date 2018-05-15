package urlshort

import (
	"net/http"
	"strings"

	"github.com/go-yaml/yaml"
)

type urlMapping struct {
	Path string
	Url  string
}

func getHandler(pathsToUrls []urlMapping, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		url := "/" + strings.TrimLeft(request.URL.Path, "/")

		for _, mapping := range pathsToUrls {
			if mapping.Path == url {
				http.RedirectHandler(mapping.Url, 301).ServeHTTP(rw, request)
			}
		}
		fallback.ServeHTTP(rw, request)
	}
}

// HANDLERS
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mappings := make([]urlMapping, len(pathsToUrls))
	i := 0
	for k, v := range pathsToUrls {
		mappings[i] = urlMapping{Path: k, Url: v}
		i++
	}
	return getHandler(mappings, fallback)
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	items := make([]urlMapping, 0)
	yaml.Unmarshal(yml, &items)

	return getHandler(items, fallback), nil
}
