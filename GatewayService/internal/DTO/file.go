package DTO

import (
	"fmt"
	"strings"
)

type (
	Item struct {
		Files []File `json:"files"`
	}

	File struct {
		Username string `json:"username"`
		Lang     string `json:"lang"`
		Data     string `json:"data"`
	}
)

func Parse(item Item) (result []File) {
	for _, file := range item.Files {
		result = append(result, file)
	}
	return
}

func ParseString(item Item) string {
	builder := strings.Builder{}
	for _, file := range item.Files {
		builder.WriteString(fmt.Sprintf("%s, %s, %s;", file.Username, file.Lang, file.Data))
	}
	return builder.String()
}
