package repl

import (
	"fmt"
	"strconv"

	"github.com/availity/av/util"
)

// Port changes the port
func (shell *Shell) Port(args []string) {
	if len(args) == 0 {
		util.LogInfo(fmt.Sprintf("%d", shell.prompt.Port))
	} else {
		port, err := strconv.Atoi(args[0])
		if err == nil {
			shell.prompt.Port = port
			shell.refreshClient()
		} else {
			util.LogError("Port must be a number")
		}
	}
}
