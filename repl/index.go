package repl

func (shell *Shell) Index(args string) {
	shell.prompt.Index = args
}
