package repl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

// Search performs a search
func (shell *Shell) Search(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	if shell.prompt.index == "" {
		return "", errors.New("index required")
	}

	service := shell.client.Search().Index(shell.prompt.index)
	for key, value := range parseTerms(args) {
		service = service.Query(elastic.NewTermQuery(key, value))
	}
	res, err := service.Do(shell.ctx)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("Time: %d ms", res.TookInMillis))
	buf.WriteString(fmt.Sprintf("Total hits: %d", res.TotalHits()))

	if res.Hits != nil {
		for _, hit := range res.Hits.Hits {
			source, err := json.Marshal(&hit.Source)
			if err != nil {
				return "", err
			}
			buf.WriteString(fmt.Sprint("ID: ", hit.Id))
			buf.WriteString(string(source))
			buf.WriteString("\n")
		}
	}
	return buf.String(), nil
}
