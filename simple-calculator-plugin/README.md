# What is this

This directory aims to confirm how calculator-plugin works
This content is based on [Go 1.8のpluginパッケージを試してみる)](https://qiita.com/qt-luigi/items/47a7913145fc747da0c7)

## How to confirm

### Build Pugin

    % task -d . build
    task: [build] cd calc ; go build -buildmode=plugin
    task: [build] go build -o app

### Checking build file

    % tree 
    .
    ├── README.md
    ├── Taskfile.yml
    ├── app
    ├── calc
    │   ├── calc.go
    │   └── calc.so
    ├── go.mod
    └── main.go

    1 directory, 7 files

### Run Plugin

    % task -d . run  
    task: [build] cd calc ; go build -buildmode=plugin
    task: [build] go build -o app
    task: [run] ./app
    init() of calc.
    pow2: 3 * 3 = 9
    add: 5 + 2 = 7
    sub: 5 - 2 = 3
