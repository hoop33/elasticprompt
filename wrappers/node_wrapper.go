package wrappers

import (
	"bytes"
	"fmt"

	"github.com/olivere/elastic"
)

type NodeWrapper struct {
	Node *elastic.NodesInfoNode
}

func NewNodeWrapper(node *elastic.NodesInfoNode) *NodeWrapper {
	return &NodeWrapper{
		Node: node,
	}
}

func (wrapper *NodeWrapper) String() string {
	var buffer bytes.Buffer

	if node := wrapper.Node; node != nil {
		// Name, version, build
		buffer.WriteString(fmt.Sprintf("%s (v%s/build %s)\n", node.Name, node.Version, node.Build))

		// Host and IP
		buffer.WriteString(fmt.Sprintf("%s/%s\n", node.Host, node.IP))

		// OS Information
		buffer.WriteString(NewOSWrapper(node.OS).String())

		// Settings
		//buffer.WriteString("  Settings:\n")
		//for key, value := range node.Settings {
		//buffer.WriteString(fmt.Sprintf("    %s: %s\n", key, value))
		//}

    buffer.WriteString("-----\n")
	}
	return buffer.String()
}
