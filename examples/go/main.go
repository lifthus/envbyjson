package main

import (
	"fmt"
	"log"
	"os"

	envbyjson "github.com/lifthus/envbyjson/go"
)

func main() {
	// loads env vars from env.json.
	err := envbyjson.Load()
	if err != nil {
		log.Fatal(err)
	}

	// loads env vars from env.json and env2.json.
	err = envbyjson.Load("env.json", "env2.json")
	if err != nil {
		log.Fatal(err)
	}

	// loads env vars from env3.json's "C" property.
	err = envbyjson.LoadProp("env3.json", "C")
	if err != nil {
		log.Fatal(err)
	}

	// load env vars from the "C" property of env3.json and env4.json.
	err = envbyjson.LoadProp("env3.json", "env4.json", "C")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("A=%s\n", os.Getenv("A"))
	fmt.Printf("B=%s\n", os.Getenv("B"))
	fmt.Printf("C=%s\n", os.Getenv("C"))
	fmt.Printf("D=%s\n", os.Getenv("D"))
	fmt.Printf("E=%s\n", os.Getenv("E"))
	fmt.Printf("F=%s\n", os.Getenv("F"))
	fmt.Printf("G=%s\n", os.Getenv("G"))
	fmt.Printf("H=%s\n", os.Getenv("H"))
}
