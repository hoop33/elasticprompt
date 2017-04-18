package repl

import (
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

// Connect connects to an Elasticsearch instance
func (shell *Shell) Connect(args []string) (string, error) {
	var url string
	if len(args) == 0 {
		url = elastic.DefaultURL
	} else {
		url = args[0]
	}

	client, err := elastic.NewClient(
		elastic.SetURL(url),
	)
	if err != nil {
		return "", err
	}

	shell.client = client
	shell.prompt.URL = url
	return fmt.Sprint("Connected to ", url), nil
}
