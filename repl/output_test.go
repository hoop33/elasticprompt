package repl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputDefaultsToText(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	assert.Equal(t, outputText, shell.prompt.output)
}
