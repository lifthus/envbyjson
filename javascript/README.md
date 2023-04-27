# envbyjson

the JavaScript project of [envbyjson](https://github.com/lifthus/envbyjson) which loads env vars from json files.

## Installation

As a library

```shell
npm install envbyjson

yarn add envbyjson
```

## Usage

```javascript
import envbyjson from "envbyjson";

  // it loads env.json file in the root directory
  envbyjson.load()
  console.log(process.env.A)

  // it loads env2.json file in the testdata directory
  envbyjson.load("testdata/env2.json")
  console.log(process.env.B)

  // now do whatever with the configuration values.
}

```
