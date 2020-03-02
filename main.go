package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	Version  = "0.01"
	Revision = "HEAD"
)

func main() {
	app := newApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "grouper"
	app.Usage = "Force grouped import path"
	app.Version = fmt.Sprintf("%s-%s", Version, Revision)
	app.Authors = []*cli.Author{{
		Name:  "tommy-sho",
		Email: "tomiokasyogo@gmail.com",
	}}
	app.Action = grouper
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "local",
			Aliases: []string{"l"},
			Usage:   "specify imports prefix beginning with this string after 3rd-party packages. especially your own organization name. comma-separated list",
		},
		&cli.BoolFlag{
			Name:    "write",
			Aliases: []string{"w"},
			Usage:   "write result source to original file instead od stdout",
		},
	}

	return app
}

func grouper(c *cli.Context) error {
	env := Env{
		Paths:       c.Args().Slice(),
		Write:       c.Bool("write"),
		LocalPrefix: c.String("local"),
	}

	return grouperMain(env)
}
