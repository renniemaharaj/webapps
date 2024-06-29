package webapps

// MetaTag represents a meta tag in the HTML head.
type MetaTag struct {
	Attribute string
	Values    []string
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
func MakeMeta(attribute string, values []string, content string) MetaTag {
	return MetaTag{
		Attribute: attribute,
		Values:    values,
		Content:   content,
	}
}

// TypicalMeta is a type for defining typical meta tag names
type TypicalMeta string

// Predefined typical meta tag names
const (
	Description TypicalMeta = "description"
	Keywords    TypicalMeta = "keywords"
	Author      TypicalMeta = "author"
	Charset     TypicalMeta = "charset"
	viewport    TypicalMeta = "viewport"
)

// AppendTypicalMeta appends a typical meta tag to the Metas struct
func (metas *Metas) AppendTypicalMeta(typicalMeta TypicalMeta, values []string, content string) {
	meta := MetaTag{
		Attribute: string(typicalMeta),
		Values:    values,
		Content:   content,
	}
	metas.MetaTags = append(metas.MetaTags, meta)
}
