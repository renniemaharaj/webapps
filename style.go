package webapps

import (
	"fmt"
	"strings"
)

type Style struct {
	Attribute string
	Value     string
}
type StyleBlock struct {
	Selection string
	Styles    []Style
}

func (styleblock StyleBlock) GenerateCSS() string {
	var sb strings.Builder

	//Write selection
	sb.WriteString(styleblock.Selection)

	//Open block
	sb.WriteRune('{')

	//Line break into block
	sb.WriteString("\n")

	for _, style := range styleblock.Styles {
		sb.WriteString(fmt.Sprintf("%v: %v;", style.Attribute, style.Value))
	}

	//Line break unto closing line
	sb.WriteString("\n")

	//Close block
	sb.WriteRune('}')
	return sb.String()
}
