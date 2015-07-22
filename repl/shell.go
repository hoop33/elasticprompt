package repl

import (
	"fmt"
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
