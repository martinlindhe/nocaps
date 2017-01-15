package main

import (
	"log"

	"github.com/martinlindhe/nocaps/lib"
)

func main() {
	err := nocaps.Perform()
	if err != nil {
		log.Println(err)
	}
}
