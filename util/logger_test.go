package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoGetsColorized(t *testing.T) {
	assert.True(t, strings.HasPrefix(ColorInfo("test"), "\x1b[0;"))
	assert.True(t, strings.HasSuffix(ColorInfo("test"), "\x1b[0m"))
}

func TestWarningGetsColorized(t *testing.T) {
	assert.True(t, strings.HasPrefix(ColorWarning("test"), "\x1b[0;"))
	assert.True(t, strings.HasSuffix(ColorWarning("test"), "\x1b[0m"))
}

func TestErrorGetsColorized(t *testing.T) {
	assert.True(t, strings.HasPrefix(ColorError("test"), "\x1b[0;"))
	assert.True(t, strings.HasSuffix(ColorError("test"), "\x1b[0m"))
}

func TestSuccess(t *testing.T) {
	assert.True(t, strings.HasPrefix(ColorSuccess("test"), "\x1b[0;"))
	assert.True(t, strings.HasSuffix(ColorSuccess("test"), "\x1b[0m"))
}
