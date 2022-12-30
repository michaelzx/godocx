package main

import (
	"github.com/michaelzx/godocx"
	"os"
	"time"
)

func main() {
	docx := godocx.New()
	// add text
	para1 := docx.AddParagraph()
	para1.AddText(time.Now().Format("2006-01-02") + "的天气真好")

	f, err := os.Create("./tmp/3.docx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = docx.Write(f)
	if err != nil {
		panic(err)
	}
}
