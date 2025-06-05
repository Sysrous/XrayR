package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/Aqr-K/XrayR/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
