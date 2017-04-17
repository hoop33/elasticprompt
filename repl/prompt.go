package repl

import (
	"bytes"
	"fmt"
)

// Prompt is the REPL prompt
type Prompt struct {
	URL   string
	Index string
}

// NewPrompt creates a new prompt with the defaults (localhost:9200)
func NewPrompt() *Prompt {
	return &Prompt{}
}

// Prompt displays the prompt
func (prompt *Prompt) Prompt() string {
	buf := bytes.NewBufferString("")

	if prompt.URL != "" {
		buf.WriteString(fmt.Sprintf("%s ", prompt.URL))
	}

	if prompt.Index != "" {
		buf.WriteString(fmt.Sprintf("(%s) ", prompt.Index))
	}

	buf.WriteString("> ")
	return buf.String()
}
