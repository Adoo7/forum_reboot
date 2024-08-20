package handlers

func LoadPage(title string) (*Page, error) {
	return &Page{Title: "Loaded Page"}, nil
}
