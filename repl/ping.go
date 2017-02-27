package repl

import (
	"github.com/hoop33/elasticprompt/util"
	"gopkg.in/olivere/elastic.v3"
)

// Ping pings the server
func (shell *Shell) Ping(args []string) {
	service := elastic.NewPingService(shell.client)
	result, _, err := service.Do()
	if err == nil {
		json, err := util.JSONString(result)
		if err == nil {
			util.LogInfo(json)
		} else {
			util.LogError(err.Error())
		}
	} else {
		util.LogError(err.Error())
	}
}
