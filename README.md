# Go Experiment

## Prerequisites

* Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  > Verify with `go version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  > Install with `export PATH=$PATH:``go env GOPATH``/bin`
  > Verify with `echo $PATH`

* Make sure Git is installed (version 2 or later).
  > Verify with `git --version`

## Build

* Clone the repo to your Go workspace:
```
cd `go env GOPATH`
go get github.com/orbs-network/go-experiment
cd src/github.com/orbs-network/go-experiment
git checkout master
```

* Build with `go install`.

* Install dependencies with `go get -t ./...`.
> TODO: make sure this phase is required here

## Run

* To run the pre-built binary:
```
cd `go env GOPATH`
./bin/go-experiment
```

* To rebuild from source and run:
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-experiment
go run *.go
```
