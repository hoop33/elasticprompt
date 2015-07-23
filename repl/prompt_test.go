package repl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptContainsDefaultHostAndPort(t *testing.T) {
	prompt := NewPrompt()
	assert.Equal(t, "localhost:9200 () > ", prompt.Prompt())
}

func TestPromptUpdatesHost(t *testing.T) {
  prompt := NewPrompt()
  prompt.Host = "foo"
  assert.Equal(t, "foo:9200 () > ", prompt.Prompt())
}

func TestPromptUpdatesPort(t *testing.T) {
  prompt := NewPrompt()
  prompt.Port = 1337
  assert.Equal(t, "localhost:1337 () > ", prompt.Prompt())
}

func TestPromptUpdatesIndex(t *testing.T) {
  prompt := NewPrompt()
  prompt.Index = "foo"
  assert.Equal(t, "localhost:9200 (foo) > ", prompt.Prompt())
}
