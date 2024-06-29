package webapps

//This function calls BlankDocument function, adds on typical meta tags and returns a new HtmlDocument.
func DefaultDocument() HtmlDocument {
	var mydocument = BlankDocument()

	var metas = &Metas{}
	var values = make([]string, 1)
	values[0] = "UTF-8"
	metas.AppendMeta(MakeMeta("charset", values, ""))
	values[0] = "viewport"
	metas.AppendMeta(MakeMeta("name", values, "width=device-width, initial-scale=1.0"))
	mydocument.Head.Metas = *metas
	return mydocument
}

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
