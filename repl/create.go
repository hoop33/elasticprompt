package repl

// Create creates . . . something
func (shell *Shell) Create(args []string) (string, error) {
	if !shell.IsConnected() {
		return "", ErrNotConnected
	}
	return "", nil
}
