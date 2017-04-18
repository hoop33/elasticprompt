package repl

// Exit exits the shell
func (shell *Shell) Exit(args []string) (string, error) {
	return shell.Quit(args)
}
