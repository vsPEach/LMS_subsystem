package DTO

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
