package repl

import (
	"fmt"
	"strconv"
)

// Prompt is the REPL prompt
type Prompt struct {
	Host  string
	Port  int
	Index string
}

// NewPrompt creates a new prompt with the defaults (localhost:9200)
func NewPrompt() *Prompt {
	return &Prompt{
		Host: "localhost",
		Port: 9200,
	}
}

// Prompt displays the prompt
func (prompt *Prompt) Prompt() string {
	return fmt.Sprint(
		prompt.Host,
		":",
		strconv.Itoa(prompt.Port),
		" (",
		prompt.Index,
		") > ",
	)
}
