package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	l "github.com/ehazlett/rivet/log"
	"github.com/ehazlett/rivet/version"
)

func init() {
	f := &l.SimpleFormatter{}
	log.SetFormatter(f)
}

func main() {
	app := cli.NewApp()
	app.Name = "rivet"
	app.Usage = "docker machine api backend"
	app.Version = version.FULL_VERSION
	app.Author = "@ehazlett"
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable Debug",
		},
	}
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
