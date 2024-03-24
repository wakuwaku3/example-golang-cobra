# example-golang-cobra

Go の CLI フレームワークである cobra を使った CLI のサンプル実装

## install

- required
  - go
  - make
  - [aqua](https://aquaproj.github.io/)

```sh
make install
```

## build

```sh
make build
```

## usage

### install completion

- for bash

```sh
make install-bash
```

## how to develop

### add sub command

```sh
cobra-cli add $command_name
```
