package repl

import "fmt"

// Index sets the index
func (shell *Shell) Index(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	if len(args) == 0 {
		return shell.prompt.index, nil
	}

	if args[0] == "*" {
		shell.prompt.index = ""
		return "index cleared", nil
	}

	shell.prompt.index = args[0]
	return fmt.Sprint("index set to ", shell.prompt.index), nil
}
