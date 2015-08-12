package repl

import (
	"github.com/hoop33/elasticprompt/util"
	"github.com/olivere/elastic"
)

func (shell *Shell) Indices(args string) {
	service := elastic.NewIndicesStatsService(shell.client)
	result, err := service.Do()
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
