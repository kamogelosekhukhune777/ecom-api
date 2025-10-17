package model

type File struct {
	BaseModel
	Name        string
	Directory   string
	Description string
	MimeType    string
}
