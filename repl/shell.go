package repl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hoop33/elasticprompt/util"
	"github.com/nemith/goline"
	"gopkg.in/olivere/elastic.v3"
)

// Shell is the REPL
type Shell struct {
	prompt *Prompt
	client *elastic.Client
}

// NewShell creates a new shell
func NewShell() *Shell {
	return &Shell{}
}

// Run runs the shell (REPL)
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
			args := strings.Split(strings.TrimSpace(line), " ")
			if len(args[0]) > 0 {
				method := reflect.ValueOf(shell).MethodByName(strings.Title(args[0]))
				if method.IsValid() {
					method.Call([]reflect.Value{reflect.ValueOf(args[1:])})
				} else {
					util.LogError(fmt.Sprint("Unknown command: '", args[0], "'"))
				}
			}
		}
	}
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

func parseTerms(args []string) map[string]string {
	terms := make(map[string]string)
	for _, pair := range args {
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
