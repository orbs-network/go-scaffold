# Types

Types and interfaces are defined with [protobuf](https://developers.google.com/protocol-buffers/) (version 3).

## Prerequisites

* Make sure [Protobuf](http://brewformulas.org/Protobuf) is installed (version 3.5 or later).

  > Install with `brew install protobuf`
  
  > Verify with `protoc --version`

* Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  
  > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
  
  > Verify with `echo $PATH`

* Make sure [Go protobuf plugin](https://developers.google.com/protocol-buffers/docs/gotutorial) is installed.
  
  > Install with `go get -u github.com/golang/protobuf/protoc-gen-go`
  
  > Verify with `protoc-gen-go` (it will wait for stdin, just close it)

* Make sure [Square goprotowrap](https://github.com/square/goprotowrap) is installed.
  
  > Install with `go get -u github.com/square/goprotowrap/cmd/protowrap`
  
  > Verify with `protowrap --version`

## Build

* Compile all `.proto` files to `.go` files from within this directory:
```
cd `go env GOPATH`
cd src/github.com/orbs-network/go-scaffold/types
rm ./**/*.pb.go
protowrap -I. --go_out `go env GOPATH`/src ./**/*.proto
```

* It is recommended to push all generated `.go` files to git as well.

## Todo

* Extract this folder (types and interface definitions) to a different repo (probably `orbs-network/orbs-spec`).

* Figure out a way to support type aliases in protobuf so signatures and such don't appear as `bytes`.

* Figure out a way to do contract tests on service interfaces because without gRPC they will not be auto generated.
