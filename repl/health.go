package repl

import "github.com/hoop33/elasticprompt/util"

// Health shows the cluster health
func (shell *Shell) Health(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	res, err := shell.client.ClusterHealth().Do(shell.ctx)
	if err != nil {
		return "", err
	}
	return util.JSONString(res)
}
