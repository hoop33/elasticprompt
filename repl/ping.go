package repl

import (
	"github.com/hoop33/elasticprompt/util"
	"github.com/olivere/elastic"
)

func (shell *Shell) Ping(args string) {
	service := elastic.NewPingService(shell.client)
	result, _, err := service.Do()
	if err == nil {
		json, err := util.JsonString(result)
		if err == nil {
			util.LogInfo(json)
		} else {
			util.LogError(err.Error())
		}
	} else {
		util.LogError(err.Error())
	}
}