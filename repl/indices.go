package repl

import "github.com/hoop33/elasticprompt/util"

// Indices gets the indices
func (shell *Shell) Indices(args []string) {
	response, err := shell.client.IndexGet().Do(shell.ctx)
	if err == nil {
		// TODO fix this so it scales better
		if len(args) > 0 && (args[0] == "-f" || args[0] == "--full") {
			json, err := util.JSONString(response)
			if err == nil {
				util.LogInfo(json)
			} else {
				util.LogError(err.Error())
			}
		}
	} else {
		util.LogError(err.Error())
	}
}
