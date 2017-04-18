package repl

import (
	"fmt"
	"strings"
)

// Output shows or sets the output type
func (shell *Shell) Output(args []string) (string, error) {
	if len(args) == 0 {
		return shell.prompt.output, nil
	}

	output := outputText
	if strings.EqualFold(args[0], outputJSON) {
		output = outputJSON
	}

	shell.prompt.output = output
	return fmt.Sprint("output set to ", shell.prompt.output), nil
}
