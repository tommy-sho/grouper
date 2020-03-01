package main

import "os"

const version = "0.0.1"

var revision = "HEAD"

func main() {
	app := newApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}