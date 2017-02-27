package repl

import "github.com/hoop33/elasticprompt/util"

func (shell *Shell) Indices(args []string) {
	response, err := shell.client.IndexGet().Do()
	if err == nil {
		// TODO fix this so it scales better
		if len(args) > 0 && (args[0] == "-f" || args[0] == "--full") {
			json, err := util.JsonString(response)
			if err == nil {
				util.LogInfo(json)
			} else {
				util.LogError(err.Error())
			}
		} else {
			for name, _ := range response {
				util.LogInfo(name)
			}
		}
	} else {
		util.LogError(err.Error())
	}
}
