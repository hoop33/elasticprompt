package repl

import (
	"bytes"
	"fmt"
)

const (
	outputText = "Text"
	outputJSON = "JSON"
)

// prompt is the REPL prompt
type prompt struct {
	url    string
	index  string
	output string
}

// newPrompt creates a new prompt
func newPrompt() *prompt {
	return &prompt{
		output: outputText,
	}
}

// prompt displays the prompt
func (prompt *prompt) Prompt() string {
	buf := bytes.NewBufferString("")

	if prompt.url != "" {
		buf.WriteString(fmt.Sprintf("%s ", prompt.url))
	}

	if prompt.index != "" {
		buf.WriteString(fmt.Sprintf("(%s) ", prompt.index))
	}

	if prompt.output == "" {
		prompt.output = outputText
	}

	buf.WriteString(fmt.Sprintf("%c ", prompt.output[0]))

	buf.WriteString("> ")
	return buf.String()
}
