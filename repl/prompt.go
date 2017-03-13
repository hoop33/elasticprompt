package repl

import "fmt"

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
	return fmt.Sprint(
		prompt.URL,
		" (",
		prompt.Index,
		") > ",
	)
}
