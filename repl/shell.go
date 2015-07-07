package repl

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/hoop33/elasticprompt/util"
	"github.com/nemith/goline"
	"github.com/olivere/elastic"
)

type Shell struct {
	prompt *Prompt
	client *elastic.Client
}

func NewShell() *Shell {
	return &Shell{}
}

func (shell *Shell) refreshClient() {
	url := fmt.Sprint("http://", shell.prompt.Host, ":", strconv.Itoa(shell.prompt.Port), "/")
	util.LogInfo(fmt.Sprint("Connecting to ", url, "..."))
	client, err := elastic.NewClient(
		elastic.SetURL(url),
	)
	if err == nil {
		shell.client = client
	} else {
		util.LogError(err.Error())
	}
}

func parseLine(line string) (string, string) {
	command := ""
	arguments := ""

	tokens := strings.SplitN(line, " ", 2)
	count := len(tokens)
	if count > 0 {
		command = strings.Title(tokens[0])
	}
	if count > 1 {
		arguments = tokens[1]
	}
	return command, arguments
}

func parseTerms(args string) map[string]string {
	terms := make(map[string]string)
	for _, pair := range strings.Split(args, "&") {
		if pair != "" {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				terms[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			} else {
				util.LogError(fmt.Sprint("Can't parse ", pair))
			}
		}
	}
	return terms
}

func (shell *Shell) Run() {
	shell.prompt = NewPrompt()
	shell.refreshClient()

	gl := goline.NewGoLine(shell.prompt)

	for {
		line, err := gl.Line()
		if err != nil {
			if err == goline.UserTerminatedError {
				return
			}
			util.LogError(err.Error())
		} else {
			fmt.Println()
			line = strings.TrimSpace(line)
			if len(line) > 0 {
				command, arguments := parseLine(line)

				method := reflect.ValueOf(shell).MethodByName(command)
				if method.IsValid() {
					method.Call([]reflect.Value{reflect.ValueOf(arguments)})
				} else {
					util.LogError(fmt.Sprint("Unknown command: '", command, "'"))
				}
			}
		}
	}
}

func (shell *Shell) Quit(args string) {
	os.Exit(0)
}

func (shell *Shell) Host(args string) {
	shell.prompt.Host = args
	shell.refreshClient()
}

func (shell *Shell) Port(args string) {
	port, err := strconv.Atoi(args)
	if err == nil {
		shell.prompt.Port = port
		shell.refreshClient()
	} else {
		util.LogError("Port must be a number")
	}
}

func (shell *Shell) Index(args string) {
	shell.prompt.Index = args
}

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
