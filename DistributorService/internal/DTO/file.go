package DTO

import (
	"fmt"
)

type Item struct {
	Files []struct {
		Username string `json:"username"`
		Lang     string `json:"lang"`
		Data     string `json:"data"`
	} `json:"files"`
}

func (i *Item) ToStringSlice() []string {
	result := make([]string, 0, 10)
	for _, file := range i.Files {
		result = append(result, fmt.Sprintf("username:%s, lang:%s, data:%s", file.Username, file.Lang, file.Data))
	}
	return result
}
