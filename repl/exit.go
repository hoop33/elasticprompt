package repl

// Exit exits the shell
func (shell *Shell) Exit(args []string) {
	shell.Quit(args)
}
