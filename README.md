# Go Experiment

## Prerequisites

* Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  > Verify with `go version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  > Install with `export PATH=$PATH:``go env GOPATH``/bin`
  > Verify with `echo $PATH`

* Make sure Git is installed (version 2 or later).
  > Verify with `git --version`

## Prerequisites for tests

* Make sure [Ginkgo](http://onsi.github.io/ginkgo/#getting-ginkgo) is installed.
  > Install with `go get github.com/onsi/ginkgo/ginkgo`
  > Verify with `ginkgo version`

## Build

* Clone the repo to your Go workspace:
```
cd `go env GOPATH`
go get github.com/orbs-network/go-experiment
cd src/github.com/orbs-network/go-experiment
git checkout master
```

* Build with `go install`

* Install dependencies with `go get -t ./...`
> TODO: make sure this phase is required here

## Run

* To run the pre-built binary (should be in path):
```
go-experiment
```

* To rebuild from source and run (this will take you to project root):
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-experiment
go run *.go
```

## Test

* Run **all** tests from project root with `ginkgo ./...`

* Another alternative runner with minimal UI is `go test ./...`

#### E2E Spec
> End to end tests (server only) checking compliance to spec

* Run the suite from project root with `ginkgo -v ./e2e/spec/`

#### E2E Stress
> End to end tests (server only) under extreme random load

* Run the suite from project root with `ginkgo -v ./e2e/stress/`

#### Service Spec
> Component test checking a single service compliance to spec

* For StateStorage, run the suite from project root with `ginkgo -v ./services/statestorage/spec`

* For VirtualMachine, run the suite from project root with `ginkgo -v  ./services/virtualmachine/spec`

* For PublicApi, run the suite from project root with `ginkgo -v  ./services/publicapi/spec`