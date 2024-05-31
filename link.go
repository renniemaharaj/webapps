package webapps

//Links of type []Link for document head
type Links struct {
	LinkTags []LinkTag
}

// LinkTag represents a link tag in the HTML head.
type LinkTag struct {
	Rel  string
	Href string
}

// Create functions for easy adding and removing link tags
func (links *Links) AppendLink(link LinkTag) {
	links.LinkTags = append(links.LinkTags, link)
}
