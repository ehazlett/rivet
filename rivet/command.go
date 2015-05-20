package main

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func run(c *cli.Context) {
	_, err := exec.LookPath("pluginhook")
	if err != nil {
		log.Fatal("unable to find pluginhook.  see https://github.com/progrium/pluginhook")
	}

}
