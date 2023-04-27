# envbyjson [![Go Report Card](https://goreportcard.com/badge/github.com/lifthus/envbyjson/go)](https://goreportcard.com/report/github.com/lifthus/envbyjson/go)

the Golang project of [envbyjson](https://github.com/lifthus/envbyjson) which loads env vars from json files.

## Installation

As a library

```shell
go get github.com/lifthus/envbyjson/go
```

## Usage

Add your application configuration to json files.

if you don't specify the file paths, envbyjson loads "env.json" in your project root directory.

```shell
{
  "A": "got A",
  "B": "got B"
}
```

in your Go app you can do something like

```go
package main

import (
    "log"
    "os"

    envbyjson "github.com/lifthus/envbyjson/go"
)

func main() {
  err := envbyjson.Load()
  if err != nil {
    log.Fatal("error loading env.json file")
  }

  // LoadProp loads properties(nested json) in "tmp" property in given json files and sets them to the env vars.
  err := envbyjson.LoadProp("env2.json", "env3.json", "tmp")
  if err != nil {
    log.Fatal("error loading configuration files")
   }


  resultA := os.Getenv("A")

  // now do whatever with the configuration values.
}
```
