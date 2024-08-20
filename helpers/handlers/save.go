package handlers

import (
	"os"
)

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("temp/"+filename, []byte("saved from save.go file"), 0600)
}
