package repl

import (
	"github.com/hoop33/elasticprompt/util"
	"github.com/olivere/elastic"
)

func (shell *Shell) Nodes(args string) {
	service := elastic.NewNodesInfoService(shell.client)
	response, err := service.Do()
	if err == nil {
		json, err := util.JsonString(response)
		if err == nil {
			util.LogInfo(json)
		} else {
			util.LogError(err.Error())
		}
	} else {
		util.LogError(err.Error())
	}
}