package repl

import "os"

// Quit quits the shell
func (shell *Shell) Quit(args []string) (string, error) {
	os.Exit(0)
	return "", nil
}
