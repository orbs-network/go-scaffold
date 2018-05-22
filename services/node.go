package services

import (
	"github.com/orbs-network/go-scaffold/services/publicapi"
	"github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/services/statestorage"
	"github.com/orbs-network/go-scaffold/utils/logger"
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

func NewNode(logger logger.Interface) Node {
	return &node{
		logger: logger,
		publicApi: publicapi.NewService(logger),
		virtualMachine: virtualmachine.NewService(logger),
		stateStorage: statestorage.NewService(logger),
		publicApiServer: publicapi.NewServer(logger),
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
