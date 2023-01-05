package main

import (
	"github.com/michaelzx/godocx"
	"os"
	"time"
)

func main() {
	docx := godocx.New()
	body := docx.Body()
	body.AddParagraph().AddText(time.Now().Format("2006-01-02") + "的天气真好")
	body.AddParagraph().AddText("size=22").Size(22)
	body.AddParagraph().AddText("color=808080").Color("808080")
	body.AddParagraph().AddText("color=ff0000").Color("ff0000")
	body.AddParagraph().AddText("size=22 and color=ff0000").Size(22).Color("ff0000")
	body.AddParagraph().AddText("size=22 and font=微软雅黑").FontEastAsia("微软雅黑")
	body.AddParagraph().AddText("----------------------------------------")

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
