package main

import (
	"github.com/michaelzx/godocx"
)

func main() {
	docx := godocx.New()
	err := docx.Save("./tmp/1.docx")
	if err != nil {
		panic(err)
	}
}
