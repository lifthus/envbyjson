# envjson [![Go Report Card](https://goreportcard.com/badge/github.com/lifthus/envjson/go)](https://goreportcard.com/report/github.com/lifthus/envjson/go)

the Golang project of [envjson](https://github.com/lifthus/envjson) which loads env vars from json files.

## Installation

As a library

```shell
go get github.com/lifthus/envjson/go
```

or if you want to use it as a bin command

go >= 1.17
```shell
go install github.com/lifthus/envjson/go@latest
```

go < 1.17
```shell
go get github.com/lifthus/envjson/go
```

## Usage

Add your application configuration to json files.

if you don't specify the file paths, envjson loads "env.json" in your project root directory.

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

    envjson "github.com/lifthus/envjson/go"
)

func main() {
  err := envjson.Load()
  if err != nil {
    log.Fatal("error loading env.json file")
  }
  
  // LoadProp loads properties(nested json) in "tmp" property in given json files and sets them to the env vars.
  err := envjson.LoadProp("env2.json", "env3.json", "tmp")
  

  resultA := os.Getenv("A")

  // now do whatever with the configuration values.
}
```
