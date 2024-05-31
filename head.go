package webapps

import (
	"fmt"
	"strings"
)

// HtmlHead represents the head section of an HTML document.
type HtmlHead struct {
	Title       string
	Description string
	Keywords    string
	Author      string
	Metas       Metas
	Links       Links
	Scripts     Scripts
}

// GenerateMarkup for HtmlHead generates the HTML for the head section.
func (head HtmlHead) GenerateMarkup() string {
	var sb strings.Builder

	//Begin markup generation of head with opening tag
	sb.WriteString("<head>\n")

	//Place title tag if title was given else default title
	if head.Title != "" {
		sb.WriteString(fmt.Sprintf("  <title>%s</title>\n", head.Title))
	} else {
		sb.WriteString("<title>webapps</title>\n")
	}

	//Generate meta tags markup
	for _, meta := range head.Metas.MetaTags {
		if meta.Content != "" {
			sb.WriteString(fmt.Sprintf("  <meta %s=\"%s\" content=\"%s\">\n", meta.Attribute, meta.Value, meta.Content))
		}
		if meta.Content == "" {
			sb.WriteString(fmt.Sprintf("  <meta %s=\"%s\">\n", meta.Attribute, meta.Value))
		}
	}

	//Generate links tags markup
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

	//End Html head markup generation with closing tag
	sb.WriteString("</head>\n")

	return sb.String()
}
