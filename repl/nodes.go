package repl

import (
	"fmt"

	"github.com/availity/av/util"
	"github.com/hoop33/elasticprompt/wrappers"
	"github.com/olivere/elastic"
)

func (shell *Shell) Nodes(args string) {
	service := elastic.NewNodesInfoService(shell.client)
	response, err := service.Do()
	if err == nil {
		util.LogInfo(fmt.Sprint("Cluster: ", response.ClusterName))
		util.LogInfo("Nodes:")
		for _, node := range response.Nodes {
			nodeWrapper := wrappers.NewNodeWrapper(node)
			util.LogInfo(nodeWrapper.String())
		}
	} else {
		util.LogError(err.Error())
	}
}
