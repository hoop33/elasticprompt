package repl

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hoop33/elasticprompt/util"
	"github.com/nemith/goline"
	"gopkg.in/olivere/elastic.v5"
)

// Shell is the REPL
type Shell struct {
	prompt *Prompt
	client *elastic.Client
	ctx    context.Context
}

// NewShell creates a new shell
func NewShell(ctx context.Context) *Shell {
	return &Shell{
		ctx: ctx,
	}
}

// Run runs the shell (REPL)
func (shell *Shell) Run() error {
	shell.prompt = NewPrompt()

	gl := goline.NewGoLine(shell.prompt)

	for {
		line, err := gl.Line()
		if err != nil {
			if err == goline.UserTerminatedError {
				return nil
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
