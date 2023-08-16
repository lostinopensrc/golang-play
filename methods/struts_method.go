package methods

import "fmt"

type Book struct {
	Name   string
	Author string
}

func (b Book) Artifact() {
	fmt.Println("Book Name:", b.Name)
	fmt.Println("Author Name:", b.Author)
}
