package services

import (
	"github.com/orbs-network/go-experiment/services/publicapi"
	"github.com/orbs-network/go-experiment/services/virtualmachine"
	"github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/utils/logger"
)

type Node interface {
	Start(stop *chan error)
	Stop()
}

type node struct {
	stop *chan error
	logger logger.Interface
	publicApi publicapi.Service
	virtualMachine virtualmachine.Service
	stateStorage statestorage.Service
	publicApiServer publicapi.Server
}

func NewNode() Node {
	logger := logger.DefaultLogger("node1")
	return &node{
		logger: logger,
		publicApi: publicapi.NewService(),
		virtualMachine: virtualmachine.NewService(),
		stateStorage: statestorage.NewService(logger),
		publicApiServer: publicapi.NewServer(),
	}
}

func (n *node) Start(stop *chan error) {
	if n.stop == nil {
		n.stop = stop
		n.stateStorage.Start(stop)
		n.virtualMachine.Start(n.stateStorage, stop)
		n.publicApi.Start(n.virtualMachine, stop)
		n.publicApiServer.Start(n.publicApi, stop)
		n.logger.Info("Node (as a whole) started")
	}
}

func (n *node) Stop() {
	if n.stop != nil {
		n.logger.Info("Node (as a whole) stopping")
		n.publicApiServer.Stop()
		n.publicApi.Stop()
		n.virtualMachine.Stop()
		n.stateStorage.Stop()
		n.stop = nil
	}
}
