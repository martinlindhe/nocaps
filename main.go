package main

import (
	"log"

	"github.com/martinlindhe/nocaps/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	revert = kingpin.Flag("revert", "Revert configuration.").Bool()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	var err error
	if *revert {
		err = nocaps.EnableCaps()
	} else {
		err = nocaps.DisableCaps()
	}
	if err != nil {
		log.Println(err)
	}
}
