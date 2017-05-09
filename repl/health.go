package repl

import (
	"fmt"

	"github.com/hoop33/elasticprompt/util"
)

// Health shows the cluster health
func (shell *Shell) Health(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	res, err := shell.client.ClusterHealth().Do(shell.ctx)
	if err != nil {
		return "", err
	}

	if shell.prompt.output == outputText {
		return fmt.Sprintf("Name: %s; Status: %s; Nodes: %d", res.ClusterName, res.Status, res.NumberOfNodes), nil
	}

	return util.JSONString(res)
}
