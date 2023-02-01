# What is this

This directory aims to confirm how terrappoi works
This content is based on [Terraform の plugin を理解する](https://qiita.com/TsuyoshiUshio@github/items/c87236e29af16927b13b)

## How to confirm

### Build Pugin

    % task -d . build
    task: [build] go build -o terrappoi .
    task: [build] go build -o ./plugin/azure_provider ./plugin/azure_provider.go

### Checking build file

    % tree 
    .
    ├── README.md
    ├── Taskfile.yml
    ├── common
    │   └── provider.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── plugin
    │   ├── azure_provider
    │   └── azure_provider.go
    └── terrappoi

    2 directories, 9 files

### Run Plugin

    % task -d . run  
    task: [build] go build -o terrappoi .
    task: [build] go build -o ./plugin/azure_provider ./plugin/azure_provider.go
    task: [run] ./terrappoi
    2023-02-01T16:20:22.384+0900 [WARN]  client: plugin configured with a nil SecureConfig
    2023-02-01T16:20:22.384+0900 [DEBUG] client: starting plugin: path=./plugin/azure_provider args=["./plugin/azure_provider"]
    2023-02-01T16:20:22.386+0900 [DEBUG] client: plugin started: path=./plugin/azure_provider pid=31963
    2023-02-01T16:20:22.386+0900 [DEBUG] client: waiting for RPC address: path=./plugin/azure_provider
    2023-02-01T16:20:22.514+0900 [DEBUG] client.azure_provider: Now Azure Provider Serving: timestamp="2023-02-01T16:20:22.514+0900"
    2023-02-01T16:20:22.514+0900 [DEBUG] client: using plugin: version=1
    2023-02-01T16:20:22.514+0900 [DEBUG] client.azure_provider: plugin address: address=/var/folders/08/m3z1dggn7f3dd8v_rw096t_w0000gn/T/plugin3683502602 network=unix timestamp="2023-02-01T16:20:22.514+0900"
    2023-02-01T16:20:22.516+0900 [DEBUG] client.azure_provider: Creating Azure Resources ...: timestamp="2023-02-01T16:20:22.516+0900"
    123
    2023-02-01T16:20:22.516+0900 [DEBUG] client.azure_provider: 2023/02/01 16:20:22 [DEBUG] plugin: plugin server: accept unix /var/folders/08/m3z1dggn7f3dd8v_rw096t_w0000gn/T/plugin3683502602: use of closed network connection
    2023-02-01T16:20:22.517+0900 [INFO]  client: plugin process exited: path=./plugin/azure_provider pid=31963
    2023-02-01T16:20:22.517+0900 [DEBUG] client: plugin exited
