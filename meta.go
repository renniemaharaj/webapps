package webapps

// MetaTag represents a meta tag in the HTML head.
type MetaTag struct {
	Attribute string
	Value     string
	Content   string
}

// type MetaTags
type Metas struct {
	MetaTags []MetaTag
}

// Create functions for easy adding and removing meta tags
func (metas *Metas) AppendMeta(metatag MetaTag) {
	metas.MetaTags = append(metas.MetaTags, metatag)
}
