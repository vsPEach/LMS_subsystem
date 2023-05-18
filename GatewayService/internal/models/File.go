package models

type (
	File struct {
		Name   string `json:"name"`
		Issuer string `json:"issuer"`
		Type   string `json:"type"`
		Data   string `json:"data"`
	}
	Item struct {
		Files []File `json:"files"`
	}
)
