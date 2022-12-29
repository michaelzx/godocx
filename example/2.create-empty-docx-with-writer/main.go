package main

import (
	"github.com/michaelzx/godocx"
	"os"
)

func main() {
	docx := godocx.New()
	f, err := os.Create("./tmp/2.docx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = docx.Write(f)
	if err != nil {
		panic(err)
	}
}
