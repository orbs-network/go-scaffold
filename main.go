package main

import (
	"github.com/orbs-network/go-scaffold/services"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"log"
)

func main() {
	stop := make(chan error, 10) //TODO: find a better way to handle this than buffered with 10 slots
	logger := logger.DefaultLogger("node1")
	node := services.NewNode(logger)
	node.Start(&stop)
	log.Fatal(<- stop)
}
