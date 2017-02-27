package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hoop33/elasticprompt/repl"
)

func main() {

	app := cli.NewApp()
	app.Name = "elasticprompt"
	app.Version = "0.0.1"
	app.Usage = "Interact with elasticsearch from the command line"
	app.Authors = []cli.Author{
		{Name: "Rob Warner", Email: "rwarner@grailbox.com"},
	}

	app.Action = func(c *cli.Context) {
		shell := repl.NewShell()
		shell.Run()
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
