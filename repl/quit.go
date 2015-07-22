package repl

import "os"

func (shell *Shell) Quit(args string) {
	os.Exit(0)
}
