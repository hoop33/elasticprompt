package repl

import (
	"strconv"

	"github.com/availity/av/util"
)

func (shell *Shell) Port(args string) {
	port, err := strconv.Atoi(args)
	if err == nil {
		shell.prompt.Port = port
		shell.refreshClient()
	} else {
		util.LogError("Port must be a number")
	}
}
