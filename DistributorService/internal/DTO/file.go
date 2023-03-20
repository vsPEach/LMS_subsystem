package DTO

type File struct {
	User      *User  `json:"user"`
	Extension string `json:"extension"`
	data      []byte `json:"data"`
}

func NewFile() *File {
	return &File{
		User:      NewUser(),
		Extension: "py",
		data:      nil,
	}
}
