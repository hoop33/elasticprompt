package repl

func (shell *Shell) Host(args string) {
	shell.prompt.Host = args
	shell.refreshClient()
}
