package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zhuharev/shop/cmd"
)

func main() {
	app := &cli.App{
		Commands: cli.Commands{
			cmd.CmdBot,
		},
	}
	app.Run(os.Args)
}
