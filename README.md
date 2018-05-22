# Go Scaffold

A simple project with 3 microservices all running in the same process for easy debugging (using protobuf as interface but without using gRPC as a transport layer).

HTTP requests are sent to *PublicApi* then forwarded for execution by *VirtualMachine* which relies on *StateStorage* to read/write state. This process involves a single node only.

The project is thoroughly tested with unit tests (100% coverage), component tests per microservice (defining its spec), E2E tests (defining the full system spec) and E2E stress tests (running the system under load).

## Prerequisites

* Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  > Verify with `go version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
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
go get github.com/orbs-network/go-scaffold
cd src/github.com/orbs-network/go-scaffold
git checkout master
```

* Install dependencies with `go get -t ./...`

* Build with `go install`

## Run

* To run the pre-built binary (should be in path):
```
go-scaffold
```

* To rebuild from source and run (this will take you to project root):
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-scaffold
go run *.go
```

## Test

* Run **all** tests from project root with `ginkgo ./...`

* Run everything **except stress** with `ginkgo --skip=Stress ./...`

* Another alternative runner with minimal UI and result caching:

  * Run **all** tests with `go test ./...` 
    > Note: E2E is flaky here for some reason
  
  * Run everything **except stress** with `go test -short ./...`
  
* Check unit test coverage with:
    ```
    go test -cover `go list ./... | grep -v /stress | grep -v /spec | grep -v /types/`
    ```

##### E2E Spec
> End to end tests (server only) checking compliance to spec

* Run the suite from project root with `ginkgo -v ./e2e/spec/`

##### E2E Stress
> End to end tests (server only) under extreme random load

* Run the suite from project root with `ginkgo -v ./e2e/stress/`

##### Service Spec
> Component test checking a single service compliance to spec

* For StateStorage, run the suite from project root with `ginkgo -v ./services/statestorage/spec`

* For VirtualMachine, run the suite from project root with `ginkgo -v  ./services/virtualmachine/spec`

* For PublicApi, run the suite from project root with `ginkgo -v  ./services/publicapi/spec`