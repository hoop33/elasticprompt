package repl

import "github.com/availity/av/util"

func (shell *Shell) Host(args []string) {
	if len(args) == 0 {
		util.LogInfo(shell.prompt.Host)
	} else {
		shell.prompt.Host = args[0]
		shell.refreshClient()
	}
}
