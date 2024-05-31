package webapps

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// HtmlDocument represents an entire HTML document.
type HtmlDocument struct {
	Language string
	Head     HtmlHead
	Body     Body
}

type Body struct {
	Elements []Element
}

// AppendChild for document body
func (body *Body) AppendChild(element Element) {
	body.Elements = append(body.Elements, element)
}

// Calls HtmlDocument.GemerateMarkup function and writes it to specified file.
// Specify .html as the extension for filename to export as html file
func (doc HtmlDocument) ExportMarkup(filename string) {
	html := doc.GenerateMarkup()
	fileName := filename
	err := os.WriteFile(fileName, []byte(html), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

// GenerateMarkup() generates the complete HTML document, beginning with webapps credits
func (doc HtmlDocument) GenerateMarkup() string {
	var sb strings.Builder

	// Add comments at the top of the document
	sb.WriteString("<!--\n")
	sb.WriteString("Credit newrennie/webapps")
	sb.WriteString("-->\n")

	//Begin html markup generation
	sb.WriteString("<!DOCTYPE html>\n")

	//Insert specified language else insert en
	if doc.Language != "" {
		sb.WriteString(fmt.Sprintf("<html lang=\"%v\">\n", doc.Language))
	} else {
		sb.WriteString("<html lang=\"en\">\n")
	}

	//Generate Document Head
	sb.WriteString(doc.Head.GenerateMarkup())

	//Begin document body markup generation
	sb.WriteString("<body>\n")

	//Let each element return their own markup generation
	for _, elem := range doc.Body.Elements {
		sb.WriteString(elem.GenerateMarkup() + "\n")
	}

	//End document body markup generation with closing tag
	sb.WriteString("</body>\n")

	//End html markup generation with closing tag
	sb.WriteString("</html>")

	return sb.String()
}
