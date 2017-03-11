package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/ympons/yax/cmd/yax/root"

	// commands
	_ "github.com/ympons/yax/cmd/yax/github"
	_ "github.com/ympons/yax/cmd/yax/init"
	_ "github.com/ympons/yax/cmd/yax/zsh"
)

func main() {
	if err := root.Command.Execute(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
