# What is this

This directory aims to confirm how helloplug works
This content is based on [go-plugin × gRPC で自作Goツールにプラグイン機構を実装する方法)](https://po3rin.com/blog/go-plug)

## How to confirm

### Build Pugin

    % task -d . build
    task: [build] go build ./cmd/helloplug/main.go
    task: [build] go build -o hello-plugin ./plugins/hello/main.go

### Checking build file

    % tree 
    .
    ├── README.md
    ├── Taskfile.yml
    ├── cmd
    │   └── helloplug
    │       └── main.go
    ├── go.mod
    ├── go.sum
    ├── hello-plugin
    ├── helloplug.go
    ├── main
    ├── plug
    │   └── plug.go
    ├── plugins
    │   └── hello
    │       └── main.go
    └── proto
        └── greeter
            ├── greeter.pb.go
            ├── greeter.proto
            └── greeter_grpc.pb.go

    7 directories, 13 files

### Run Plugin

    % task -d . run  
    task: [build] go build ./cmd/helloplug/main.go
    task: [build] go build -o hello-plugin ./plugins/hello/main.go
    task: [run] export GREETER_PLUGIN="./hello-plugin"
    task: [run] ./main
    hello gopher
