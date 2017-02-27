package repl

import "github.com/availity/av/util"

// Index sets the index
func (shell *Shell) Index(args []string) {
	if len(args) == 0 {
		util.LogInfo(shell.prompt.Index)
	} else {
		shell.prompt.Index = args[0]
	}
}
