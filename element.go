package webapps

import (
	"fmt"
	"strings"
)

// Element represents a generic HTML element.
type Element struct {
	Tag        string
	Attributes []Attribute
	Children   []Element
}

// ApendChild() for Elements appends children
func (parent *Element) AppendChild(child Element) {
	parent.Children = append(parent.Children, child)
}

// AppendAttribute for Elements appends attributes using pointers
func (element *Element) AppendAttribute(attribute Attribute) {
	element.Attributes = append(element.Attributes, attribute)
}

// GenerateMarkup() generates the HTML for an Element, including its children.
func (elem Element) GenerateMarkup() string {
	var sb strings.Builder

	// Start the opening tag
	sb.WriteString(fmt.Sprintf("<%s", elem.Tag))

	//Gnerate all attributes
	for _, attribute := range elem.Attributes {
		sb.WriteString(fmt.Sprintf(" %v=\"%v\"", attribute.Name, attribute.Value))
	}

	//End opening tag
	sb.WriteString(">")

	// Recursively generate children elements
	for _, child := range elem.Children {
		sb.WriteString(child.GenerateMarkup())
	}

	// Close the tag
	sb.WriteString(fmt.Sprintf("</%s>", elem.Tag))

	return sb.String()
}
