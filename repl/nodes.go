package repl

import "github.com/hoop33/elasticprompt/util"

// Nodes gets the nodes
func (shell *Shell) Nodes(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	res, err := shell.client.NodesInfo().Do(shell.ctx)
	if err != nil {
		return "", err
	}
	return util.JSONString(res)
}
