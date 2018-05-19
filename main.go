package main

import (
	"github.com/orbs-network/go-experiment/services/publicapi"
	"github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/services/virtualmachine"
	"log"
)

func main() {
	// vars
	stop := make(chan error)

	// create services
	publicApi := publicapi.NewService()
	virtualMachine := virtualmachine.NewService()
	stateStorage := statestorage.NewService()

	// start services
	stateStorage.Start(&stop)
	virtualMachine.Start(&stateStorage, &stop)
	publicApi.Start(&virtualMachine, &stop)

	// create and start servers
	publicapi.NewServer().Start(&publicApi, &stop)

	// wait until stop
	log.Fatal(<- stop)
}
