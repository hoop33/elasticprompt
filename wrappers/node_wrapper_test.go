package wrappers

import (
	"testing"

	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
)

func TestNilNodeReturnsEmpty(t *testing.T) {
	assert.Equal(t, "", NewNodeWrapper(nil).String())
}

func TestBasicNodeInformationIsPresent(t *testing.T) {
	doc := `My Cool Node (v1.2.3/build abcdef)
example/8.8.8.8
-----
`
	node := &elastic.NodesInfoNode{
		Name:    "My Cool Node",
		Host:    "example",
		IP:      "8.8.8.8",
		Version: "1.2.3",
		Build:   "abcdef",
	}
	assert.Equal(t, doc, NewNodeWrapper(node).String())
}
