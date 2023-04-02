package DTO

type File struct {
	User string `json:"username"`
	Lang string `json:"lang"`
	Data []byte `json:"data"`
}

type Files []File
