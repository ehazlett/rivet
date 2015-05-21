package main

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/rivet/api"
)

func run(c *cli.Context) {
	// check for pluginhook
	if _, err := exec.LookPath("pluginhook"); err != nil {
		log.Fatal("unable to find pluginhook.  see https://github.com/progrium/pluginhook")
	}

	listenAddr := c.String("listen-addr")
	hooksPath := c.String("hooks-path")

	cfg := &api.ApiConfig{
		ListenAddr: listenAddr,
		HooksPath:  hooksPath,
	}
	a := api.NewApi(cfg)

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
