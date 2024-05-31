package webapps

// NewDocument creates and returns an HtmlDocument skeleton.
func NewDocument() HtmlDocument {
	return HtmlDocument{
		Head: HtmlHead{
			Title:   "",
			Metas:   Metas{},
			Links:   Links{},
			Scripts: Scripts{},
		},
		Body: Body{},
	}
}

//CreateAttribute will create attributes in map[name]value
func CreateAttribute(name, value string) Attribute {
	return Attribute{
		Name:  name,
		Value: value,
	}
}

// CreateElement function creates a new HTML element with the specified tag, attributes: class id and onclick, and content.
// func CreateElement(tag, class, id, onclick, content string) Element {

// 	return Element{Tag: tag, Class: class, ID: id, OnClick: onclick, Content: content}
// }
// CreateElement creates an element with specified tag, class, id, onclick, and content
func CreateElement(tag, class, id, onclick, content string) Element {
	attrs := []Attribute{}

	if class != "" {
		attrs = append(attrs, Attribute{Name: "class", Value: class})
	}
	if id != "" {
		attrs = append(attrs, Attribute{Name: "id", Value: id})
	}
	if onclick != "" {
		attrs = append(attrs, Attribute{Name: "onclick", Value: onclick})
	}
	if content != "" {
		attrs = append(attrs, Attribute{Name: "content", Value: onclick})
	}

	return CreateElementByAttributes(tag, &attrs)
}

//This will create and element of type specificed, and take a map for attributes.
func CreateElementByAttributes(tag string, attributes *[]Attribute) Element {
	var element Element = Element{}
	element.Tag = tag
	element.Attributes = *attributes
	return element
}

//Returns a simple meta tag [name content]
func MakeMeta(attribute, value, content string) MetaTag {
	return MetaTag{
		Attribute: attribute,
		Value:     value,
		Content:   content,
	}
}

//Returns a new meta tag
func MakeLink(rel, href string) LinkTag {
	return LinkTag{
		Rel:  rel,
		Href: href,
	}
}

//Returns a new script tag
func MakeScript(src string, async, deferring bool, inline string) Script {
	return Script{
		Src:    src,
		Async:  async,
		Defer:  deferring,
		Inline: inline,
	}
}

//Calls NewDocument() and adds standard meta tags then returns modified HtmlDocument
func DefaultDocument() HtmlDocument {
	var mydocument = NewDocument()

	var metas = &Metas{}
	metas.AppendMeta(MakeMeta("charset", "UTF-8", ""))
	metas.AppendMeta(MakeMeta("name", "viewport", "width=device-width, initial-scale=1.0"))
	mydocument.Head.Metas = *metas
	return mydocument
}
