# What is this

This directory aims to confirm how go-plugin works
This content is based on [Goメモ-220 (標準ライブラリのpluginパッケージで簡易プラグイン処理)](https://devlights.hatenablog.com/entry/2022/07/04/073000)

## How to confirm

### Build Pugin

    % task -d . build
    task: [build] cd lib ; go build -buildmode=plugin
    task: [build] go build -o app

### Checking build file

    % tree 
    .
    ├── Makefile
    ├── Taskfile.yml
    ├── app
    ├── go.mod
    ├── lib
    │   ├── lib.go
    │   ├── lib.so
    │   └── pkg
    │       └── strs
    │           └── upper.go
    └── main.go

    3 directories, 8 files

### Run Plugin

    % task -d . run  
    task: [build] cd lib ; go build -buildmode=plugin
    task: [build] go build -o app
    task: [run] ./app
    HELLO WORLD
