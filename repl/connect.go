package repl

import (
	"fmt"

	"github.com/availity/av/util"
	"gopkg.in/olivere/elastic.v5"
)

// Connect connects to an Elasticsearch instance
func (shell *Shell) Connect(args []string) {
	if len(args) == 0 {
		util.LogError("You must specify the URL to the Elasticsearch host")
		return
	}

	// TODO add port if not present
	url := args[0]
	util.LogInfo(fmt.Sprint("Connecting to ", url, "..."))
	client, err := elastic.NewClient(
		elastic.SetURL(url),
	)
	if err != nil {
		util.LogError(err.Error())
	} else {
		shell.client = client
		shell.prompt.URL = url
	}
}
