package main

import (
	"github.com/orbs-network/go-scaffold/services"
	"log"
)

func main() {
	stop := make(chan error, 10) //TODO: find a better way to handle this than buffered with 10 slots
	node := services.NewNode()
	node.Start(&stop)
	log.Fatal(<- stop)
}
