package repl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputShouldDefaultToText(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetTextForNil(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output(nil)
	assert.Nil(t, err)
	assert.Equal(t, outputText, msg)
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetTextForEmpty(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{})
	assert.Nil(t, err)
	assert.Equal(t, outputText, msg)
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetTextForBadString(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{"bad"})
	assert.Nil(t, err)
	assert.Equal(t, "output set to Text", msg)
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetTextFortext(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{"text"})
	assert.Nil(t, err)
	assert.Equal(t, "output set to Text", msg)
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetTextForTEXT(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{"TEXT"})
	assert.Nil(t, err)
	assert.Equal(t, "output set to Text", msg)
	assert.Equal(t, outputText, shell.prompt.output)
}

func TestOutputShouldSetJSONForjson(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{"json"})
	assert.Nil(t, err)
	assert.Equal(t, "output set to JSON", msg)
	assert.Equal(t, outputJSON, shell.prompt.output)
}

func TestOutputShouldSetJSONForJSON(t *testing.T) {
	shell := &Shell{}
	shell.prompt = newPrompt()
	msg, err := shell.Output([]string{"JSON"})
	assert.Nil(t, err)
	assert.Equal(t, "output set to JSON", msg)
	assert.Equal(t, outputJSON, shell.prompt.output)
}
