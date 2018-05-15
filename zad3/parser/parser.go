package parser

import (
	"encoding/json"
	"os"

	"gopher/zad3/models"
)

type Provider = func(string) (models.Story, bool)

func CreateProvider(filePath string) Provider {
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var stories map[string]models.Story
	json.NewDecoder(reader).Decode(&stories)

	return func(arcName string) (models.Story, bool) {
		story, ok := stories[arcName]
		return story, ok
	}
}
