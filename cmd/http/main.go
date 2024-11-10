package main

import log "github.com/sirupsen/logrus"

func main() {
	err := start(":9000")
	if err != nil {
		log.Fatalf("[start]: %v", err)
	}
}
