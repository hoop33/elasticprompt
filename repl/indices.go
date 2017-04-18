package repl

import "github.com/hoop33/elasticprompt/util"

// Indices gets the indices
func (shell *Shell) Indices(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	if len(args) == 0 {
		args = []string{"_all"}
	}

	res, err := shell.client.IndexGet(args...).Do(shell.ctx)
	if err != nil {
		return "", err
	}
	return util.JSONString(res)
}
