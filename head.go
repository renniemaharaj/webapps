package webapps

import (
	"fmt"
	"strings"
)

// HtmlHead struct represents the entire head section of a document.
type HtmlHead struct {
	Title       string
	Description string
	Keywords    string
	Author      string
	Metas       Metas
	Links       Links
	Scripts     Scripts
}

// This function generates the markup for an HtmlHead struct.
func (head HtmlHead) GenerateMarkup() string {
	var sb strings.Builder

	//Begin markup generation of head with opening tag.
	sb.WriteString("<head>\n")

	//Place title if it was given, otherwise place a default one.
	if head.Title != "" {
		sb.WriteString(fmt.Sprintf("  <title>%s</title>\n", head.Title))
	} else {
		sb.WriteString("<title>webapps</title>\n")
	}

	//Generate all other meta tags.
	for _, meta := range head.Metas.MetaTags {
		if meta.Content != "" {
			sb.WriteString(fmt.Sprintf("  <meta %s=\"%s\" content=\"%s\">\n", meta.Attribute, strings.Join(meta.Values, ","), meta.Content))
		}
		if meta.Content == "" {
			sb.WriteString(fmt.Sprintf("  <meta %s=\"%s\">\n", meta.Attribute, strings.Join(meta.Values, ",")))
		}
	}

	//Generate link tags markup.
	for _, link := range head.Links.LinkTags {
		sb.WriteString(fmt.Sprintf("  <link rel=\"%s\" href=\"%s\">\n", link.Rel, link.Href))
	}

	//Generate script tags markup
	for _, script := range head.Scripts.ScriptTags {
		if script.Inline != "" {
			sb.WriteString(fmt.Sprintf("  <script>%s</script>\n", script.Inline))
		} else {
			asyncAttr := ""
			deferAttr := ""
			if script.Async {
				asyncAttr = " async"
			}
			if script.Defer {
				deferAttr = " defer"
			}
			sb.WriteString(fmt.Sprintf("  <script src=\"%s\"%s%s></script>\n", script.Src, asyncAttr, deferAttr))
		}
	}

	//End Html head markup generation with closing tag.
	sb.WriteString("</head>\n")

	return sb.String()
}
