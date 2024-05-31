package webapps

//Scripts tag of []Script for document head
type Scripts struct {
	ScriptTags []Script
}

// Script represents a script tag in the HTML head.
type Script struct {
	Src    string
	Async  bool
	Defer  bool
	Inline string
}

// Create functions for easy adding and removing script tags
func (scripts *Scripts) AppendScript(script Script) {
	scripts.ScriptTags = append(scripts.ScriptTags, script)
}
