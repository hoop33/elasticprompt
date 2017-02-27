package repl

import "os"

// Quit quits the shell
func (shell *Shell) Quit(args []string) {
	os.Exit(0)
}
