package repl

import (
	"fmt"

	"github.com/hoop33/elasticprompt/util"
)

func (shell *Shell) Nodes(args []string) {
	response, err := shell.client.NodesInfo().Do()
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
			util.LogInfo(fmt.Sprintf("cluster name: %s", response.ClusterName))
			for _, node := range response.Nodes {
				util.LogInfo(node.Name)
			}
		}
	} else {
		util.LogError(err.Error())
	}
}
