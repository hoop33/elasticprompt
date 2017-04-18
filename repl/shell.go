package repl

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/hoop33/elasticprompt/util"
	"github.com/nemith/goline"
	"gopkg.in/olivere/elastic.v5"
)

// ErrNotConnected means the shell is not connected to Elasticsearch
var ErrNotConnected = errors.New("Not connected")

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

// IsConnected returns whether the client is connected
func (shell *Shell) IsConnected() bool {
	return shell.client != nil
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
					res := method.Call([]reflect.Value{reflect.ValueOf(args[1:])})
					output := res[0].Interface().(string)
					err := res[1].Interface()
					if err != nil {
						util.LogError(err.(error).Error())
					} else {
						util.LogInfo(output)
					}
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
