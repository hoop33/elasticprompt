package repl

import (
	"github.com/hoop33/elasticprompt/util"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Ping pings the server
func (shell *Shell) Ping(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	res, _, err := elastic.NewPingService(shell.client).Do(shell.ctx)
	if err != nil {
		return "", err
	}

	return util.JSONString(res)
}
