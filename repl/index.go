package repl

import "fmt"

// Index sets the index
func (shell *Shell) Index(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}

	if len(args) == 0 {
		return shell.prompt.Index, nil
	}

	if args[0] == "*" {
		shell.prompt.Index = ""
		return "Index cleared", nil
	}

	shell.prompt.Index = args[0]
	return fmt.Sprint("Index set to ", shell.prompt.Index), nil
}
