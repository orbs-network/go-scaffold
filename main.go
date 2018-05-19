package main

import (
	"log"
)

func main() {
	stop := make(chan error)
	node := NewNode()
	node.Start(&stop)
	log.Fatal(<- stop)
}
