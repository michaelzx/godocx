package main

import (
	"github.com/michaelzx/godocx"
	"os"
	"time"
)

func main() {
	docx := godocx.New()
	body := docx.Body()
	// add text
	para1 := body.AddParagraph()
	para1.AddText(time.Now().Format("2006-01-02") + "的天气真好")
	para1.AddText("test font size").Size(22)
	para1.AddText("test color").Color("808080")

	para2 := body.AddParagraph()
	para2.AddText("test font size and color").Size(22).Color("ff0000")

	f, err := os.Create("./tmp/4.docx")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = docx.Write(f)
	if err != nil {
		panic(err)
	}
}
