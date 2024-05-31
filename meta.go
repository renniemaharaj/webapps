package webapps

// MetaTag represents a meta tag in the HTML head.
type MetaTag struct {
	Attribute string
	Value     string
	Content   string
}

// Type Metas represents the entire meta section of a document head
type Metas struct {
	MetaTags []MetaTag
}

// This function appends a meta tag to the metas struct of a document head.
func (metas *Metas) AppendMeta(metatag MetaTag) {
	metas.MetaTags = append(metas.MetaTags, metatag)
}

//Returns a meta tag <meta attribute="value" content="content">. Omit content if necessary
func MakeMeta(attribute, value, content string) MetaTag {
	return MetaTag{
		Attribute: attribute,
		Value:     value,
		Content:   content,
	}
}
