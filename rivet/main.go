package main

import (
	"os"
	"path/filepath"

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
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %s", err)
	}

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
		cli.StringFlag{
			Name:  "listen-addr, l",
			Usage: "Listen address",
			Value: ":8080",
		},
		cli.StringFlag{
			Name:  "hooks-path, p",
			Usage: "Path to hooks directory",
			Value: filepath.Join(wd, "hooks"),
		},
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
