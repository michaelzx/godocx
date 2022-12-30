package tpl

import _ "embed"

var (
	//go:embed rels.xml
	rels string
	//go:embed docProps/app.xml
	docPropsApp string
	//go:embed docProps/app.xml
	docPropsCore string
	//go:embed Content_Types.xml
	contentTypes string
	//go:embed word/theme/theme1.xml
	wordThemeTheme1 string
	//go:embed word/styles.xml
	wordStyles string
)
