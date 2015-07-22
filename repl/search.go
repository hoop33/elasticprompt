package repl

import (
	"encoding/json"
	"fmt"

	"github.com/availity/av/util"
	"github.com/olivere/elastic"
)

func (shell *Shell) Search(args string) {
	service := shell.client.Search().Index(shell.prompt.Index)
	for key, value := range parseTerms(args) {
		service = service.Query(elastic.NewTermQuery(key, value))
	}
	searchResult, err := service.Do()
	if err == nil {
		util.LogInfo(fmt.Sprintf("Time: %d ms", searchResult.TookInMillis))
		util.LogInfo(fmt.Sprintf("Total hits: %d", searchResult.TotalHits()))

		if searchResult.Hits != nil {
			for _, hit := range searchResult.Hits.Hits {
				source, err := json.Marshal(&hit.Source)
				if err == nil {
					util.LogInfo(fmt.Sprint("ID: ", hit.Id))
					util.LogInfo(string(source))
					fmt.Println()
				} else {
					util.LogError(err.Error())
				}
			}
		}
	} else {
		util.LogError(err.Error())
	}
}
