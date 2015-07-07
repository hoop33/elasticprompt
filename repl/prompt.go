package repl

import (
	"fmt"
	"strconv"
)

type Prompt struct {
	Host  string
	Port  int
	Index string
}

func NewPrompt() *Prompt {
	return &Prompt{
		Host: "localhost",
		Port: 9200,
	}
}

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
