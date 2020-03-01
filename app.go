package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "grouper"
	app.Usage = "Force grouped import path"
	app.Version = fmt.Sprintf("%s-%s", version, revision)
	app.Authors = []*cli.Author{{
		Name:  "tommy-sho",
		Email: "tomiokasyogo@gmail.com",
	}}
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "local",
			Usage: "specify imports prefix beginning with this string after 3rd-party packages. especially your own organization name. comma-separated list",
		},
		&cli.BoolFlag{
			Name:  "write",
			Usage: "write result source to original file instead od stdout",
		},
	}

	return app
}
